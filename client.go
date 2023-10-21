package goQL

import (
	"encoding/json"
	"io"
)


// GoQLClient is an interface for a GraphQL HTTP client.
//
// It provides two methods, `Query()` and `Mutation()`, for making GraphQL queries and mutations, respectively.
type GoQLClient interface {
	Query(q Query) GoQLClientRunner
	Mutation(m Query) GoQLClientRunner
}

// GoQL is a struct that implements the GoQLClient interface.
//
// It has the following fields:
//
// * `url`: The URL of the GraphQL endpoint.
// * `timeout`: The timeout in seconds for GraphQL requests.
// * `retry`: The number of times to retry a GraphQL request if it fails.
// * `body`: The body of the GraphQL request.
// * `method`: The HTTP method for the GraphQL request.
// * `headers`: The HTTP headers for the GraphQL request.
type GoQL struct {
	url     string
	timeout int
	retry   int
	body    io.Reader
	method  string
	headers map[string]string
}

// Query implements the GoQLClient interface for `GoQL`.
//
// It takes a `Query` struct as input and returns a `GoQLClientRunner` struct.
//
// The `Query` struct represents a GraphQL query.
//
// The `GoQLClientRunner` struct allows you to run the GraphQL query and get the results.
func (g *GoQL) Query(query Query) GoQLClientRunner {
	body, err := json.Marshal(query)
	if err != nil {
		L.Fatal("Query error:", err)
	}
	return &Runner{
		method:     g.method,
		url:        g.url,
		body:       body,
		headers:    g.headers,
		timeoutSec: g.timeout,
	}
}

// Mutation implements the GoQLClient interface for `GoQL`.
//
// It takes a `Query` struct as input and returns a `GoQLClientRunner` struct.
//
// The `Query` struct represents a GraphQL mutation.
//
// The `GoQLClientRunner` struct allows you to run the GraphQL mutation and get the results.
func (g *GoQL) Mutation(query Query) GoQLClientRunner {
	body, err := json.Marshal(query)
	if err != nil {
		L.Fatal("Query error:", err)
	}
	return &Runner{
		method:     g.method,
		url:        g.url,
		body:       body,
		headers:    g.headers,
		timeoutSec: g.timeout,
	}
}
