package example

const URL_GQL = ""                                                            // Your test url
const METHOD = "POST"                                                         // coverage-ignore
const QUERY = `query {
	post(id: 1) {
	  id
	  title
	  body
	}
  }` // coverage-ignore

const MUTATION = `mutation ($input: CreatePostInput!) {
	createPost(input: $input) {
	  id
	  title
	  body
	}
}` // coverage-ignore

var MUTATION_INPUT = struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}{
	Title: "New Post Title",
	Body:  "Some interesting content.",
} // coverage-ignore

const HTTPTimeoutSec = 2 // coverage-ignore
