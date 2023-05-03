package goQL

import (
	"encoding/json"
	"io"
)

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

// Query implements GoQLClient
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

// Mutation implements GoQLClient
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
