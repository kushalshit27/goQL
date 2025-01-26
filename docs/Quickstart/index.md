---
layout: default
title: Quickstart
nav_order: 4
has_children: true
---

## ⚡️ Quickstart

```go
package main

import "github.com/kushalshit27/goQL"

func main() {
	const GRAPHQL_URL = "{{ graphql URL }}"

	// This is a reusable client that can be used to make multiple queries 
	client := goQL.New().
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

	response := client.Query(query).Run(context.TODO())

	log.Println(response)
}
```

For advanced config see configuration section.

