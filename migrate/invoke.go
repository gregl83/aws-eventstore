package main

import (
	"github.com/aws/aws-lambda-go/events"
	"log"
)

func main() {
	//request events.APIGatewayProxyRequest
	request := events.APIGatewayProxyRequest{

	};
	response, err := Handler(request)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}