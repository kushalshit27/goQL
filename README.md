# goQL
goQL is a graphql HTTP client,with :battery: included


<!-- GETTING STARTED -->
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

	query := goQL.Query{
		Query: QUERY,
	}

	response := c.Query(query).Run(context.TODO())

	log.Println(response)
}
```


## Usage

_For more examples, please refer to the [example](example)_



## Roadmap

See the [open issues](https://github.com/github_username/repo_name/issues) for a list of proposed features (and known issues).



## License

Distributed under the MIT License. See `LICENSE` for more information.
