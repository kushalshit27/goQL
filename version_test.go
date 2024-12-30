package goQL_test

import (
	"testing"

	"github.com/kushalshit27/goQL"
)

func TestVersion(t *testing.T) {
	expected := "1.0.0"
	if goQL.Version != expected {
		t.Errorf("expected %s, got %s", expected, goQL.Version)
	}
}
