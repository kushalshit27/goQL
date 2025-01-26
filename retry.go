package main

import (
	"context"
	"github.com/kushalshit27/goQL"
	"log"
)

func main() {
	const GRAPHQL_URL = "{{ graphql URL }}"

	// This is a reusable client that can be used to make multiple queries
	gqlClient := goQL.New().
		SetURL(GRAPHQL_URL).
		SetMethod("POST").
		Build()

	query := goQL.Query{
		Query: `query {
                        post {
                            id
                            title
                            body
                        }
                    }`,
	}

	response := gqlClient.Query(query).Run(context.TODO())

	log.Println(response)
}
