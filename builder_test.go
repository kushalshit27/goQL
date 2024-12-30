package goQL_test

import (
	"testing"

	"github.com/kushalshit27/goQL"
)

func TestGoQLBuilder(t *testing.T) {
	builder := goQL.New().
		SetURL("http://example.com").
		SetMethod("POST").
		SetTimeout(10).
		SetRetry(3)

	client := builder.Build()

	if client == nil {
		t.Errorf("expected client to be initialized, got nil")
	}
}
