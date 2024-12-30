package goQL_test

import (
	"testing"
	"time"

	"github.com/kushalshit27/goQL"
)

func TestExponentialBackoff(t *testing.T) {
	backoff := goQL.ExponentialBackoff(time.Second)
	start := time.Now()
	backoff(1)
	duration := time.Since(start)

	if duration < 2*time.Second {
		t.Errorf("expected at least 2 seconds, got %v", duration)
	}
}

func TestLinearBackoff(t *testing.T) {
	backoff := goQL.LinearBackoff(time.Second)
	start := time.Now()
	backoff(1)
	duration := time.Since(start)

	if duration < 2*time.Second {
		t.Errorf("expected at least 2 seconds, got %v", duration)
	}
}
