package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	_ "github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/aws_s3"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {

	// fixme sourceUrl and database connection url

	//m, err := migrate.New(
	//	"github://mattes:personal-access-token@mattes/migrate_test",
	//	"postgres://localhost:5432/database?sslmode=enable")
	//m.Steps(2)

	fmt.Println("Received body: ", request.Body)

	return Response{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
