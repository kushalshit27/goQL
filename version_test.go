package goQL

import (
	"testing"
)

func TestVersion(t *testing.T) {
	expected := "1.0.0"
	if Version != expected {
		t.Errorf("expected %s, got %s", expected, Version)
	}
}
