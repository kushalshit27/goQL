---
layout: default
title: CHANGELOG
nav_order: 5
---

# CHANGELOG
{: .fs-9 }

## v1.2.0
{: .fs-6 }

[//]: # ({: .new })
> New (v1.2.0)
>
>   - `goQL.Query` and `goQL.Mutation` support more features like:
>   - `Debug()`: Enables debug logging for the query execution, providing additional information about the request and response.
>   - `RetryAttempts(..)`: Configures the client to retry the query in case of errors.
>   - `RetryBackoff(...)`: Defines a linear backoff strategy for retries. This means the waiting time between retries increases linearly with each attempt.
>   - `RetryOn(func(err error) bool { ... })`: Defines a custom retry logic based on the encountered error. For example, the provided function checks if the error message contains the word `timeout`, indicating a connection timeout scenario. In such cases, the query will be retried.
>   - `RetryAllowStatus(func(status int) bool { ... })`: Sets conditions for retrying on specific HTTP status codes returned by the database. The provided function allows retries for specific status codes.


## v1.0.0
{: .fs-6 }

[//]: # ({: .new })
> New (v1.0.0)
>
>   - GraphQL query builder with support for query.
>   - GraphQL mutation builder with support for mutation.

