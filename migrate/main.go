package main

import (
	"fmt"

	"github.com/gregl83/aws-eventstore/infrastructure/database"
	"github.com/gregl83/aws-eventstore/infrastructure/filestore"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	_ "github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/aws_s3"
	"github.com/golang-migrate/migrate"

	"log"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	// fixme sourceUrl and database connection url
	m, err := migrate.New(
		filestore.GetStorageURL("event-store/migrations"),
		database.GetConnectionURL("root", "password", "127.0.0.1", "events", "3306"))

	if err != nil {
		log.Fatal(err)

		return Response{Body: request.Body, StatusCode: 200}, nil
	}

	m.Up()

	fmt.Println("Received body: ", request.Body)

	return Response{Body: request.Body, StatusCode: 200}, nil
}

// main starts lambda using Handler
func main() {
	lambda.Start(Handler)
}


