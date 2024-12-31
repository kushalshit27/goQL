package goQL

import (
	"testing"
)

func TestGoQLBuilder(t *testing.T) {
	builder := New().
		SetURL("http://example.com").
		SetMethod("POST").
		SetTimeout(10).
		SetRetry(3)

	client := builder.Build()

	if client == nil {
		t.Errorf("expected client to be initialized, got nil")
	}
}
