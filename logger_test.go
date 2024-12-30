package goQL_test

import (
	"testing"

	"github.com/kushalshit27/goQL"
)

func TestLoggerInitialization(t *testing.T) {
	if goQL.L == nil {
		t.Errorf("expected logger to be initialized, got nil")
	}
}
