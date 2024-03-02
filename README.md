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

### ✨  Features:

``goQL.Query`` and ``goQL.Mutation`` supports more features like:

- `Debug()`: Enables debug logging for the query execution, providing additional information about the request and response.([example](example/query.go))
- `RetryAttempts(..)`: Configures the client to retry the query in case of errors.([example](example/query.go))
- `RetryBackoff(...)`: Defines a linear backoff strategy for retries. This means the waiting time between retries increases linearly with each attempt.([example](example/query.go))
- `RetryOn(func(err error) bool { ... })`: Defines a custom retry logic based on the encountered error. For example, The provided function checks if the error message contains the word "timeout", indicating a connection timeout scenario. In such cases, the query will be retried.([example](example/query.go))
- `RetryAllowStatus(func(status int) bool { ... })`: Sets conditions for retrying on specific HTTP status codes returned by the database. The provided function allows retries for status.([example](example/query.go))


## ⚙️ Installation
```
go get -u github.com/kushalshit27/goQL
```

## Usage

_For more examples, please refer to the [example](example)_



## Roadmap

See the [open issues](https://github.com/kushalshit27/goQL/issues) for a list of proposed features (and known issues).



## License

Distributed under the MIT License. See `LICENSE` for more information.
