package goQL

import (
	"errors"
	"net/http"
	"testing"
	"time"
)

func TestDefaultTransientError(t *testing.T) {
	err := errors.New("transient error")
	if DefaultTransientError(err) {
		t.Errorf("expected false, got true")
	}
}

func TestDefaultAllowRetryStatus(t *testing.T) {
	allowRetryStatus := DefaultAllowRetryStatus()
	if !allowRetryStatus(500) {
		t.Errorf("expected true for status 500, got false")
	}
	if allowRetryStatus(200) {
		t.Errorf("expected false for status 200, got true")
	}
}

func TestEnableDebugLog(t *testing.T) {
	opts := defaultRetryOptions
	enableDebugLog(nil)(&opts)
	if opts.DebugLog {
		t.Errorf("expected DebugLog to be false, got true")
	}

	flag := true
	enableDebugLog(&flag)(&opts)
	if !opts.DebugLog {
		t.Errorf("expected DebugLog to be true, got false")
	}
}

func TestRetryAttempts(t *testing.T) {
	opts := defaultRetryOptions
	retryAttempts(5)(&opts)
	if opts.Attempts != 5 {
		t.Errorf("expected Attempts to be 5, got %v", opts.Attempts)
	}
}

func TestRetryBackoff(t *testing.T) {
	opts := defaultRetryOptions
	backoff := LinearBackoff(time.Second)
	retryBackoff(backoff)(&opts)
	if opts.Backoff == nil {
		t.Errorf("expected Backoff to be set, got nil")
	}
}

func TestRetryOnFunc(t *testing.T) {
	opts := defaultRetryOptions
	retryOn := func(err error) bool { return true }
	retryOnFunc(retryOn)(&opts)
	if opts.RetryOn == nil {
		t.Errorf("expected RetryOn to be set, got nil")
	}
}

func TestAllowRetryStatusFunc(t *testing.T) {
	opts := defaultRetryOptions
	allowStatus := func(status int) bool { return status == http.StatusOK }
	allowRetryStatusFunc(allowStatus)(&opts)
	if opts.AllowRetryStatus == nil {
		t.Errorf("expected AllowRetryStatus to be set, got nil")
	}
}
