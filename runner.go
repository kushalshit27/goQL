package goQL

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

type GoQLClientRunner interface {
	Run(ctx context.Context) interface{}
	RawReq() GoQLClientRunner
	RawRes() GoQLClientRunner
}

type Runner struct {
	method     string
	url        string
	body       []byte
	headers    map[string]string
	rawReq     bool
	rawRes     bool
	timeoutSec int
}

func (r *Runner) RawReq() GoQLClientRunner {
	r.rawReq = true
	return r
}

func (r *Runner) RawRes() GoQLClientRunner {
	r.rawRes = true
	return r
}

func (r *Runner) Run(ctx context.Context) interface{} {
	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(r.timeoutSec))
	defer cancel()
	req, err := http.NewRequest(r.method, r.url, bytes.NewBuffer(r.body))
	if err != nil {
		L.Panic("ERROR from NewRequest", err)
	}
	for h, hv := range r.headers {
		req.Header.Set(h, hv)
	}

	req = req.WithContext(ctx)

	if r.rawReq {
		r.printRawRequest(req)
	}

	/*
		resp, err := http.DefaultClient.Do(req)
	*/
	resp, err := doWithRetry(ctx, req,
		retryAttempts(5),
		retryBackoff(linearBackoff(time.Second)),
		retryOnFunc(func(err error) bool {
			return strings.Contains(err.Error(), "timeout")
		}),
		allowRetryStatusFunc(func(i int) bool { return i >= 400 && i < 500 }),
	)
	if err != nil {
		L.Panic("ERROR from http client", err)
	}
	defer resp.Body.Close()

	if r.rawRes {
		r.printRawResponse(resp)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		L.Panic("ERROR from ReadAll body", err)
	}

	var response interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		L.Panic("ERROR from Unmarshal response", err)
	}

	return response
}

func (r *Runner) printRawRequest(req *http.Request) {
	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		L.Fatal(err)
	}

	L.Printf("----------------:REQUEST:----------------\n%s", string(reqDump))
}

func (r *Runner) printRawResponse(resp *http.Response) {
	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		L.Fatal(err)
	}

	L.Printf("\n----------------:RESPONSE:----------------\n%s", string(respDump))
}

// RetryAttempts configures the number of retry attempts
func retryAttempts(attempts int) func(*retryOptions) {
	return func(b *retryOptions) {
		b.Attempts = attempts
	}
}

// RetryBackoff configures the backoff duration between retries
func retryBackoff(backoff backoffFunc) func(*retryOptions) {
	return func(b *retryOptions) {
		b.Backoff = backoff
	}
}

// RetryOnFunc configures a function to decide retry based on error
func retryOnFunc(fn func(error) bool) func(*retryOptions) {
	return func(b *retryOptions) {
		b.RetryOn = fn
	}
}

// AllowRetryStatusFunc configures a function to allow retry based on status code
func allowRetryStatusFunc(fn func(int) bool) func(*retryOptions) {
	return func(b *retryOptions) {
		b.AllowRetryStatus = fn
	}
}

func doWithRetry(ctx context.Context, req *http.Request, opts ...retryOption) (*http.Response, error) {
	// Apply options
	options := defaultRetryOptions
	for _, opt := range opts {
		opt(&options)
	}

	client := &http.Client{}

	for attempt := 0; attempt < options.Attempts; attempt++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		resp, err := client.Do(req)

		if err == nil {
			// Successful request
			return resp, nil
		}

		if !options.RetryOn(err) {
			// Error is not eligible for retry
			return nil, err
		}

		if resp != nil && !options.AllowRetryStatus(resp.StatusCode) {
			// Status code is not eligible for retry
			return resp, err
		}

		// Backoff before retrying
		options.Backoff(attempt)
	}

	err := fmt.Errorf("exceeded max attempts")
	// Exceeded max attempts
	return nil, err

}

// RetryOption defines a function to configure retry behavior
type retryOption func(*retryOptions)

// RetryOptions holds configurable retry parameters
type retryOptions struct {
	Attempts         int              `json:"attempts"`
	Backoff          backoffFunc      `json:"backoff"`
	RetryOn          func(error) bool `json:"retryOn"`
	AllowRetryStatus func(int) bool   `json:"allowRetryStatus"`
}

var defaultRetryOptions = retryOptions{
	Attempts:         3,
	Backoff:          exponentialBackoff(time.Second),
	RetryOn:          isTransientError,
	AllowRetryStatus: func(i int) bool { return i >= 500 && i < 600 },
}

// IsTransientError checks if the error is transient
func isTransientError(err error) bool {
	// Implement your logic to identify transient errors based on error types
	// or specific error messages
	return false
}

// BackoffFunc determines the delay between retries.
type backoffFunc func(attempt int)

func exponentialBackoff(base time.Duration) backoffFunc {
	return func(attempt int) {
		duration := time.Duration(math.Pow(2, float64(attempt))) * base
		time.Sleep(duration)
	}
}

func linearBackoff(duration time.Duration) backoffFunc {
	return func(attempt int) {
		time.Sleep(duration + time.Duration(attempt))
	}
}
