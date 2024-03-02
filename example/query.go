package example

import (
	"context"
	"log"
	"strings"
	"time"

	goQL "github.com/kushalshit27/goQL"
)

func Query() {
	log.Println("running query example")
	c := goQL.New().
		SetURL(URL_GQL).
		SetMethod("POST").
		SetHeader("User-Agent", "go-QL-client").
		SetHeader("Content-Type", "application/json").
		Build()

	/** Test Query **/
	query := goQL.Query{
		Query: QUERY,
	}

	res, err := c.Query(query).
		Debug().
		RetryAttempts(2).
		RetryBackoff(goQL.LinearBackoff(time.Second)).
		RetryOn(func(err error) bool {
			return strings.Contains(err.Error(), "timeout")
		}).
		RetryAllowStatus(func(status int) bool { return status >= 400 && status < 500 }).
		Run(context.TODO())
	if err != nil {
		log.Println("Client error:", err)
		return
	}

	log.Println("Client Response", res)

}
