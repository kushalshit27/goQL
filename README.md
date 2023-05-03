# goQL
goQL is a graphql HTTP client,with :battery: included

![goql-logo](https://user-images.githubusercontent.com/43465488/222173261-efc1e3b2-c569-4254-84e4-79831370680c.png)

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


## Usage

_For more examples, please refer to the [example](example)_



## Roadmap

See the [open issues](https://github.com/kushalshit27/goQL/issues) for a list of proposed features (and known issues).



## License

Distributed under the MIT License. See `LICENSE` for more information.
