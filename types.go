package goQL

type Query struct {
	Query     string `json:"query"`
	Variables `json:"variables,omitempty"`
}
type Variables struct {
	Filter interface{} `json:"filter,omitempty"`
	Input  interface{} `json:"input,omitempty"`
}
