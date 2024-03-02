package goQL

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"time"
)

// BackoffFunc determines the delay between retries.
type backoffFunc func(attempt int)

func ExponentialBackoff(base time.Duration) backoffFunc {
	return func(attempt int) {
		duration := time.Duration(math.Pow(2, float64(attempt))) * base
		time.Sleep(duration)
	}
}

func LinearBackoff(duration time.Duration) backoffFunc {
	return func(attempt int) {
		time.Sleep(duration + time.Duration(attempt))
	}
}

// DefaultTransientError checks if the error is transient
func DefaultTransientError(err error) bool {
	// Implement your logic to identify transient errors based on error types
	// or specific error messages
	return false
}

func DefaultAllowRetryStatus() func(int) bool {
	return func(i int) bool {
		return i >= 500 && i < 600
	}
}

// RetryOptions holds configurable retry parameters
type retryOptions struct {
	Attempts         int              `json:"attempts"`
	Backoff          backoffFunc      `json:"backoff"`
	RetryOn          func(error) bool `json:"retryOn"`
	AllowRetryStatus func(int) bool   `json:"allowRetryStatus"`
	DebugLog         bool             `json:"debugLog"`
}

type options struct {
	retryOptions
}

// RetryOption defines a function to configure retry behavior
type retryOption func(*retryOptions)

var defaultRetryOptions = retryOptions{
	Attempts:         3,
	Backoff:          LinearBackoff(time.Second),
	RetryOn:          DefaultTransientError,
	AllowRetryStatus: DefaultAllowRetryStatus(),
	DebugLog:         false,
}

func doWithRetry(ctx context.Context, req *http.Request, reqTimeout time.Duration, opts ...retryOption) (*http.Response, error) {
	// Apply options
	options := defaultRetryOptions
	for _, opt := range opts {
		opt(&options)
	}

	client := &http.Client{}

	for attempt := 0; attempt < options.Attempts; attempt++ {

		rCtx, cancel := context.WithTimeout(ctx, time.Second*reqTimeout)
		defer cancel()

		select {
		case <-rCtx.Done():
			return nil, rCtx.Err()
		default:
		}

		if options.DebugLog {
			L.Println("[attempt  count]", (attempt + 1))
		}

		req = req.WithContext(rCtx)

		resp, err := client.Do(req)

		if options.DebugLog && resp != nil {
			L.Println("[http status]", resp.StatusCode)
		}

		if options.DebugLog && err != nil {
			L.Println("[http error]", err.Error())
		}

		if err == nil {
			// Successful request
			return resp, nil
		}

		fmt.Println("options.RetryOn(err)", options.RetryOn)

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

// enableDebugLog configures the debug logs
func enableDebugLog(flag *bool) func(*retryOptions) {
	return func(options *retryOptions) {
		if flag != nil {
			options.DebugLog = *flag
		}
	}
}

// RetryAttempts configures the number of retry attempts
func retryAttempts(attempts int) func(*retryOptions) {
	return func(options *retryOptions) {
		if attempts != 0 {
			options.Attempts = attempts
		}
	}
}

// RetryBackoff configures the backoff duration between retries
func retryBackoff(backoff backoffFunc) func(*retryOptions) {
	return func(options *retryOptions) {
		if backoff != nil {
			options.Backoff = backoff
		}
	}
}

// RetryOnFunc configures a function to decide retry based on error
func retryOnFunc(fn func(error) bool) func(*retryOptions) {
	return func(options *retryOptions) {
		if fn != nil {
			options.RetryOn = fn
		}
	}
}

// AllowRetryStatusFunc configures a function to allow retry based on status code
func allowRetryStatusFunc(fn func(int) bool) func(*retryOptions) {
	return func(options *retryOptions) {
		if fn != nil {
			options.AllowRetryStatus = fn
		}
	}
}
