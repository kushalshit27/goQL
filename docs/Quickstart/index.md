---
layout: default
title: Quickstart
nav_order: 4
has_children: true
---

## Table of contents
{: .no_toc .text-delta }

1. TOC 
{:toc}

## ⚡️ Quickstart

```go
package main

import "github.com/kushalshit27/goQL"

func main() {
   	const GRAPHQL_URL = "{{ graphql URL }}"
	
	c := goQL.New().
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

	response := c.Query(query).Run(context.TODO())

	log.Println(response)
}
```

For advanced config see configuration section.

