
---
layout: default
title:  Configuration
parent:  Quickstart
description: "Instructions to Configuration of GoQL"
---

## ✨ Configuration
- ``goQL.Query`` and ``goQL.Mutation`` supports more features like:

- `Debug()`: Enables debug logging for the query execution, providing additional information about the request and response.([example](example/query.go))
- `RetryAttempts(..)`: Configures the client to retry the query in case of errors.([example](example/query.go))

- `RetryBackoff(...)`: Defines a linear backoff strategy for retries. This means the waiting time between retries increases linearly with each attempt.([example](example/query.go))
- `RetryOn(func(err error) bool { ... })`: Defines a custom retry logic based on the encountered error. For example, The provided function checks if the error message contains the word "timeout", indicating a connection timeout scenario. In such cases, the query will be retried.([example](example/query.go))
- `RetryAllowStatus(func(status int) bool { ... })`: Sets conditions for retrying on specific HTTP status codes returned by the database. The provided function allows retries for status.([example](example/query.go))

