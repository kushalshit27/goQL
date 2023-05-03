package example

const URL_GQL = "" // Your test url
const METHOD = "POST"
const QUERY = `query {
	post(id: 1) {
	  id
	  title
	  body
	}
  }`

const MUTATION = `mutation ($input: CreatePostInput!) {
	createPost(input: $input) {
	  id
	  title
	  body
	}
}`

var MUTATION_INPUT = struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}{
	Title: "New Post Title",
	Body:  "Some interesting content.",
}

const HTTPTimeoutSec = 2
