package goQL

import (
	"testing"
)

func TestQueryStruct(t *testing.T) {
	query := Query{
		Query: "query { test }",
		Variables: Variables{
			Filter: "filter",
			Input:  "input",
		},
	}

	if query.Query != "query { test }" {
		t.Errorf("expected query to be 'query { test }', got %s", query.Query)
	}
	if query.Variables.Filter != "filter" {
		t.Errorf("expected filter to be 'filter', got %v", query.Variables.Filter)
	}
	if query.Variables.Input != "input" {
		t.Errorf("expected input to be 'input', got %v", query.Variables.Input)
	}
}
