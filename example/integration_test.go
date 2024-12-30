package example

import (
	"context"
	"github.com/kushalshit27/goQL"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestQuery_Integration(t *testing.T) {
	// Set up a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected method POST, got %v", r.Method)
		}
		if r.URL.Path != "/graphql" {
			t.Errorf("expected URL path /graphql, got %v", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{"post":{"id":1,"title":"Test Post","body":"This is a test post"}}}`))
	}))
	defer server.Close()

	// Create a client instance
	c := goQL.New().
		SetURL(server.URL+"/graphql").
		SetMethod("POST").
		SetHeader("User-Agent", "go-QL-client").
		SetHeader("Content-Type", "application/json").
		Build()

	query := goQL.Query{
		Query: `{"query":"query { post(id: 1) { id title body } }"}`,
	}

	res, err := c.Query(query).
		Debug().
		RetryAttempts(2).
		RetryBackoff(goQL.LinearBackoff(time.Second)).
		RetryOn(func(err error) bool {
			return strings.Contains(err.Error(), "timeout")
		}).
		RetryAllowStatus(func(status int) bool { return status >= 400 && status < 500 }).
		Run(context.TODO())
	if err != nil {
		log.Println("Client error:", err)
		return
	}

	// Verify the response
	data, ok := res["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected data field in response, got %v", res)
	}

	post, ok := data["post"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected post field in data, got %v", data)
	}

	if id, ok := post["id"].(float64); !ok || id != 1 {
		t.Errorf("expected post id 1, got %v", post["id"])
	}
	if title, ok := post["title"].(string); !ok || title != "Test Post" {
		t.Errorf("expected post title 'Test Post', got %v", post["title"])
	}
	if body, ok := post["body"].(string); !ok || body != "This is a test post" {
		t.Errorf("expected post body 'This is a test post', got %v", post["body"])
	}
}

func TestMutation_Integration(t *testing.T) {
	// Set up a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected method POST, got %v", r.Method)
		}
		if r.URL.Path != "/graphql" {
			t.Errorf("expected URL path /graphql, got %v", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"data":{"createPost":{"id":1,"title":"Test Post","body":"This is a test post"}}}`))
	}))
	defer server.Close()

	// Create a client instance
	c := goQL.New().
		SetURL(server.URL+"/graphql").
		SetMethod("POST").
		SetHeader("User-Agent", "go-QL-client").
		SetHeader("Content-Type", "application/json").
		Build()

	query := goQL.Query{
		Query: `{"query":"mutation { createPost(title: \"Test Post\", body: \"This is a test post\") { id title body } }"}`,
	}

	res, err := c.Mutation(query).
		Run(context.TODO())

	if err != nil {
		log.Println("Client error:", err)
		return
	}

	// Verify the response
	data, ok := res["data"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected data field in response, got %v", res)
	}

	post, ok := data["createPost"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected createPost field in data, got %v", data)
	}

	if id, ok := post["id"].(float64); !ok || id != 1 {
		t.Errorf("expected post id 1, got %v", post["id"])
	}

}
