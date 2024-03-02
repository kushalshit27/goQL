package example

import (
	"context"

	"log"

	goQL "github.com/kushalshit27/goQL"
)

func Mutation() {
	log.Println("running mutation example")
	c := goQL.New().
		SetURL(URL_GQL).
		SetMethod("POST").
		SetHeader("User-Agent", "go-QL-client").
		SetHeader("Content-Type", "application/json").
		Build()

	/** Test Mutation **/
	query := goQL.Query{
		Query: MUTATION,
		Variables: goQL.Variables{
			Input: MUTATION_INPUT,
		},
	}

	res, err := c.Mutation(query).
		RawReq().
		RawRes().
		Run(context.TODO())

	if err != nil {
		log.Println("Client error:", err)
		return
	}

	log.Println("Client Response", res)

}
