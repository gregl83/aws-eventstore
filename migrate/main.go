package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	//"github.com/aws/aws-lambda-go/lambda"

	_ "github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/aws_s3"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/golang-migrate/migrate"
	"log"
	"github.com/aws/aws-lambda-go/lambda"
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
		"file://migrations",
		getConnectionURL("root", "password", "127.0.0.1", "events", "3306"))

	if err != nil {
		log.Fatal(err)

		return Response{Body: request.Body, StatusCode: 200}, nil
	}

	m.Up()

	fmt.Println("Received body: ", request.Body)

	return Response{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}

// getConnectionUrl returns a formatted connection URL
func getConnectionURL(username, password, host, database, port string) string {
	return fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database)
}
