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

	// This is a reusable client
	client := goQL.New().
		SetURL(GRAPHQL_URL).
		SetMethod("POST").
		Build()

	q:= `query {
		post {
		  id
		  title
		  body
		}
            }`

	query := goQL.Query{
		Query: q,
	}

	response := client.Query(query).Run(context.TODO())

	log.Println(response)
}
```

For advanced config see configuration section.

