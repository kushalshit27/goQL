package example

import (
	"context"
	"log"

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

	c.Query(query).
		RawReq().
		RawRes().
		Run(context.TODO())

}
