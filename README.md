# goQL
goQL is a graphql HTTP client,with :battery: included

<p align="center">
  <img width="250" height="200" src="goql_icon.png">
</p>

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

See the [open issues](https://github.com/kushalshit27/goQL/issues) for a list of proposed features (and known issues).



## License

Distributed under the MIT License. See `LICENSE` for more information.
