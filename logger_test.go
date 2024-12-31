package goQL

import (
	"testing"
)

func TestLoggerInitialization(t *testing.T) {
	if L == nil {
		t.Errorf("expected logger to be initialized, got nil")
	}
}
