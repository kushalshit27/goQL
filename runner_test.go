package goQL

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner_Run(t *testing.T) {
	r := &Runner{
		method:     "GET",
		url:        "http://example.com",
		headers:    map[string]string{"Content-Type": "application/json"},
		timeoutSec: 10,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := r.Run(ctx)
	if err == nil {
		t.Errorf("expected error for invalid GraphQL response, got nil")
	}
}

func TestRunner_RetryBackoff(t *testing.T) {
	r := &Runner{}
	r = r.RetryBackoff(LinearBackoff(time.Second)).(*Runner)
	if r.options.Backoff == nil {
		t.Errorf("expected Backoff to be set, got nil")
	}
}

func TestRunner_Debug(t *testing.T) {
	r := &Runner{}
	r = r.Debug().(*Runner)
	if !r.debug {
		t.Errorf("expected debug to be true, got %v", r.debug)
	}
}

func TestRunner_RawReq(t *testing.T) {
	r := &Runner{}
	r = r.RawReq().(*Runner)
	if !r.rawReq {
		t.Errorf("expected rawReq to be true, got %v", r.rawReq)
	}
}

func TestRunner_RawRes(t *testing.T) {
	r := &Runner{}
	r = r.RawRes().(*Runner)
	if !r.rawRes {
		t.Errorf("expected rawRes to be true, got %v", r.rawRes)
	}
}

func TestRunner_RetryAttempts(t *testing.T) {
	r := &Runner{}
	r = r.RetryAttempts(3).(*Runner)
	if r.options.Attempts != 3 {
		t.Errorf("expected Attempts to be 3, got %v", r.options.Attempts)
	}
}

func TestRunner_RetryOn(t *testing.T) {
	retryOn := func(err error) bool { return true }
	r := &Runner{}
	r = r.RetryOn(retryOn).(*Runner)
	if r.options.RetryOn == nil {
		t.Errorf("expected RetryOn to be set, got nil")
	}
}

func TestRunner_RetryAllowStatus(t *testing.T) {
	allowStatus := func(status int) bool { return status == http.StatusOK }
	r := &Runner{}
	r = r.RetryAllowStatus(allowStatus).(*Runner)
	if r.options.AllowRetryStatus == nil {
		t.Errorf("expected AllowRetryStatus to be set, got nil")
	}
}

func TestRunner_SetMethod(t *testing.T) {
	r := &Runner{}
	r.method = "POST"
	if r.method != "POST" {
		t.Errorf("expected method to be POST, got %v", r.method)
	}
}

func TestRunner_SetURL(t *testing.T) {
	r := &Runner{}
	r.url = "http://example.com"
	if r.url != "http://example.com" {
		t.Errorf("expected url to be http://example.com, got %v", r.url)
	}
}

func TestRunner_SetBody(t *testing.T) {
	r := &Runner{}
	r.body = []byte(`{"key":"value"}`)
	if string(r.body) != `{"key":"value"}` {
		t.Errorf("expected body to be {\"key\":\"value\"}, got %v", string(r.body))
	}
}

func TestRunner_SetHeaders(t *testing.T) {
	r := &Runner{}
	r.headers = map[string]string{"Content-Type": "application/json"}
	if r.headers["Content-Type"] != "application/json" {
		t.Errorf("expected Content-Type to be application/json, got %v", r.headers["Content-Type"])
	}
}

func TestRunner_SetTimeoutSec(t *testing.T) {
	r := &Runner{}
	r.timeoutSec = 10
	if r.timeoutSec != 10 {
		t.Errorf("expected timeoutSec to be 10, got %v", r.timeoutSec)
	}
}

func TestRunner_printRawRequest(t *testing.T) {
	r := &Runner{
		method:  "GET",
		url:     "http://example.com",
		headers: map[string]string{"Content-Type": "application/json"},
		body:    []byte(`{"key":"value"}`),
	}

	req, err := http.NewRequest(r.method, r.url, bytes.NewBuffer(r.body))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for h, hv := range r.headers {
		req.Header.Set(h, hv)
	}

	r.printRawRequest(req)
}

func TestRunner_printRawResponse(t *testing.T) {
	r := &Runner{}

	resp := httptest.NewRecorder()
	resp.WriteHeader(http.StatusOK)
	resp.Write([]byte(`{"data":"value"}`))

	r.printRawResponse(resp.Result())
}

func TestRunner_Run_NewRequestError(t *testing.T) {
	r := &Runner{
		method: "GET",
		url:    "http://[::1]:namedport",
	}

	ctx := context.Background()
	_, err := r.Run(ctx)
	if err == nil {
		t.Errorf("expected error for invalid URL, got nil")
	}
}

func TestRunner_Run_DoWithRetryError(t *testing.T) {
	r := &Runner{
		method:     "GET",
		url:        "http://example.com",
		timeoutSec: 1,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	_, err := r.Run(ctx)
	if err == nil {
		t.Errorf("expected error for context timeout, got nil")
	}
}

func TestRunner_Run_ReadAllError(t *testing.T) {
	r := &Runner{
		method: "GET",
		url:    "http://example.com",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.(http.Flusher).Flush()
		time.Sleep(time.Second)
	}))
	defer server.Close()

	r.url = server.URL

	ctx := context.Background()
	_, err := r.Run(ctx)
	if err == nil {
		t.Errorf("expected error for read all, got nil")
	}
}

func TestRunner_Run_UnmarshalError(t *testing.T) {
	r := &Runner{
		method: "GET",
		url:    "http://example.com",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`invalid json`))
	}))
	defer server.Close()

	r.url = server.URL

	ctx := context.Background()
	_, err := r.Run(ctx)
	if err == nil {
		t.Errorf("expected error for unmarshal, got nil")
	}
}

func TestRunner_Run_GraphQLError(t *testing.T) {
	r := &Runner{
		method: "GET",
		url:    "http://example.com",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"errors":[{"message":"GraphQL error"}]}`))
	}))
	defer server.Close()

	r.url = server.URL

	ctx := context.Background()
	_, err := r.Run(ctx)
	if err == nil {
		t.Errorf("expected GraphQL error, got nil")
	}
}
