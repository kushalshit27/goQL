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


type GoQL struct {
	url     string
	timeout int
	retry   int
	body    io.Reader
	method  string
	headers map[string]string
}


func (g *GoQL) Query(query Query) GoQLClientRunner {
	body, err := json.Marshal(query)
	if err != nil {
		L.Fatal("[Query] Query error:", err)
	}
	return &Runner{
		method:     g.method,
		url:        g.url,
		body:       body,
		headers:    g.headers,
		timeoutSec: g.timeout,
	}
}


func (g *GoQL) Mutation(query Query) GoQLClientRunner {
	body, err := json.Marshal(query)
	if err != nil {
		L.Fatal("[Mutation] Query error:", err)
	}
	return &Runner{
		method:     g.method,
		url:        g.url,
		body:       body,
		headers:    g.headers,
		timeoutSec: g.timeout,
	}
}
