package goQL

import (
	"io"
)

// Version of current goql package
const Version = "1.0.0"

type GoQLBuilder interface {
	SetURL(url string) GoQLBuilder
	SetHeader(name, value string) GoQLBuilder
	SetPayload(r io.Reader) GoQLBuilder
	SetMethod(method string) GoQLBuilder
	Build() GoQLClient
}

type goQLBuilder struct {
	url     string
	timeout int // second
	retry   int
	body    io.Reader
	method  string
	headers map[string]string
	version string
}

func New() GoQLBuilder {
	return &goQLBuilder{
		headers: make(map[string]string, 0),
		version: Version,
	}
}

func (g *goQLBuilder) SetURL(url string) GoQLBuilder {
	g.url = url

	return g
}

func (r *goQLBuilder) SetMethod(method string) GoQLBuilder {
	r.method = method
	return r
}

func (g *goQLBuilder) SetHeader(name, value string) GoQLBuilder {
	_, found := g.headers[name]

	if !found {
		g.headers[name] = value
	}

	return g
}

func (g *goQLBuilder) SetPayload(r io.Reader) GoQLBuilder {
	g.body = r

	return g
}

func (g *goQLBuilder) setDefault() {

	// set default method
	if g.method == "" {
		g.method = "POST"
	}

	// set default header
	_, found := g.headers["User-Agent"]
	if !found {
		g.SetHeader("User-Agent", "go-QL-client")
	}

	_, found = g.headers["Content-Type"]
	if !found {
		g.SetHeader("Content-Type", "application/json")
	}

	// set default timeout
	if g.timeout == 0 {
		g.timeout = 3
	}
}

func (g *goQLBuilder) Build() GoQLClient {

	g.setDefault()

	return &GoQL{
		url:     g.url,
		timeout: g.timeout,
		retry:   g.retry,
		body:    g.body,
		method:  g.method,
		headers: g.headers,
	}

}
