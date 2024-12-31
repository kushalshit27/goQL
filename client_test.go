package goQL

import (
	"context"
	"testing"
)

func TestGoQLClientQuery(t *testing.T) {
	client := &GoQL{
		url: "http://example.com",
	}

	query := Query{
		Query: "query { test }",
	}

	_, err := client.Query(query).Run(context.Background())
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func TestGoQLClientMutation(t *testing.T) {
	client := &GoQL{
		url: "http://example.com",
	}

	query := Query{
		Query: "mutation { test }",
	}

	_, err := client.Mutation(query).Run(context.Background())
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func TestGoQLClientQueryWithAllOptions(t *testing.T) {
	client := &GoQL{
		url: "http://example.com",
	}

	query := Query{
		Query: "query { test }",
	}

	_, err := client.Query(query).
		Debug().
		RetryAttempts(3).
		RetryBackoff(LinearBackoff(1)).
		RetryOn(func(err error) bool { return true }).
		RetryAllowStatus(func(status int) bool { return true }).
		Run(context.Background())
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}

func TestGoQLClientMutationWithAllOptions(t *testing.T) {
	client := &GoQL{
		url: "http://example.com",
	}

	query := Query{
		Query: "mutation { test }",
	}

	_, err := client.Mutation(query).
		Debug().
		RetryAttempts(3).
		RetryBackoff(LinearBackoff(1)).
		RetryOn(func(err error) bool { return true }).
		RetryAllowStatus(func(status int) bool { return true }).
		Run(context.Background())
	if err == nil {
		t.Errorf("expected an error, got nil")
	}
}
