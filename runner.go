package goQL

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"time"
)

type GoQLClientRunner interface {
	Run(ctx context.Context) (GoQLClientResponse, *GoQLClientError)
	Debug() GoQLClientRunner
	RawReq() GoQLClientRunner
	RawRes() GoQLClientRunner
	RetryAttempts(count int) GoQLClientRunner
	RetryBackoff(backoff backoffFunc) GoQLClientRunner
	RetryOn(fn func(err error) bool) GoQLClientRunner
	RetryAllowStatus(fn func(int) bool) GoQLClientRunner
}

type GoQLClientResponse map[string]interface{}
type GoQLClientError struct {
	error
	Description string
}

type Runner struct {
	method     string
	url        string
	body       []byte
	headers    map[string]string
	rawReq     bool
	rawRes     bool
	timeoutSec int
	options    options
	debug      bool
}

func (r *Runner) Debug() GoQLClientRunner {
	r.debug = true
	// r.RawReq()
	// r.RawRes()
	return r
}

func (r *Runner) RawReq() GoQLClientRunner {
	r.rawReq = true
	return r
}

func (r *Runner) RawRes() GoQLClientRunner {
	r.rawRes = true
	return r
}

func (r *Runner) RetryAttempts(count int) GoQLClientRunner {
	r.options.Attempts = count
	return r
}

func (r *Runner) RetryBackoff(backoff backoffFunc) GoQLClientRunner {
	r.options.Backoff = backoff
	return r
}

func (r *Runner) RetryOn(fn func(err error) bool) GoQLClientRunner {
	r.options.RetryOn = fn
	return r
}

func (r *Runner) RetryAllowStatus(fn func(int) bool) GoQLClientRunner {
	r.options.AllowRetryStatus = fn
	return r
}

func (r *Runner) Run(ctx context.Context) (GoQLClientResponse, *GoQLClientError) {

	req, err := http.NewRequest(r.method, r.url, bytes.NewBuffer(r.body))
	if err != nil {
		L.Println("ERROR from NewRequest", err)
		return nil, &GoQLClientError{
			error:       err,
			Description: "ERROR from NewRequest",
		}
	}
	for h, hv := range r.headers {
		req.Header.Set(h, hv)
	}

	if r.rawReq {
		r.printRawRequest(req)
	}

	//req = req.WithContext(ctx)

	resp, err := doWithRetry(ctx, req,
		time.Duration(r.timeoutSec),
		retryAttempts(r.options.Attempts),
		retryBackoff(r.options.Backoff),
		retryOnFunc(r.options.RetryOn),
		allowRetryStatusFunc(r.options.AllowRetryStatus),
		enableDebugLog(&r.debug),
	)

	if err != nil {
		L.Println("ERROR from http client", err)
		return nil, &GoQLClientError{
			error:       err,
			Description: "ERROR from http client",
		}
	}
	defer resp.Body.Close()

	if r.rawRes {
		r.printRawResponse(resp)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		L.Println("ERROR from ReadAll body", err)
		return nil, &GoQLClientError{
			error:       err,
			Description: "ERROR from ReadAll body",
		}
	}

	// Check for GraphQL errors
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err == nil {
		if errors, ok := response["errors"]; ok {
			for _, err := range errors.([]interface{}) {
				errMsg := fmt.Sprintf("GraphQL error: message: %s\n", err.(map[string]interface{})["message"])
				return nil, &GoQLClientError{
					Description: errMsg,
				}

			}

		}
	} else {
		L.Println("ERROR from Unmarshal response", err)
		return nil, &GoQLClientError{
			error:       err,
			Description: "ERROR from Unmarshal response",
		}
	}

	return response, nil
}

func (r *Runner) printRawRequest(req *http.Request) {
	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		L.Fatal(err)
	}

	L.Printf("\n----------------:REQUEST:----------------\n%s", string(reqDump))
}

func (r *Runner) printRawResponse(resp *http.Response) {
	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		L.Fatal(err)
	}

	L.Printf("\n----------------:RESPONSE:----------------\n%s", string(respDump))
}
