package main

import (
	"log"

	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func Test(t *testing.T) {
	// fixme

	request := events.APIGatewayProxyRequest{
		// todo - set params
	}

	response, err := Handler(request)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}