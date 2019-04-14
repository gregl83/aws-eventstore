package main

import (
	"os"
	"log"
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/aws_s3"
	"github.com/golang-migrate/migrate"

	"github.com/gregl83/aws-eventstore/adapters"
	"github.com/gregl83/aws-eventstore/infrastructure/database"
	"github.com/gregl83/aws-eventstore/infrastructure/filestore"
)

var (
	S3_MIGRATIONS_PATH = os.Getenv("S3_MIGRATIONS_PATH")
	AURORA_HOST = os.Getenv("AURORA_HOST")
	AURORA_PORT = os.Getenv("AURORA_PORT")
	AURORA_DATABASE = os.Getenv("AURORA_DATABASE")
	AURORA_USERNAME = os.Getenv("AURORA_USERNAME")
	AURORA_PASSWORD string
)

// Event payload processed by lambda handler
type Event struct {
	Name string `json:"name"`
}

// Response sent by lambda handler for a given event
type Response struct {
	Body       string `json:"body"`
	StatusCode int    `json:"statusCode"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, event Event) (Response, error) {
	// fixme sourceUrl and database connection url
	m, err := migrate.New(
		filestore.GetStorageURL(S3_MIGRATIONS_PATH),
		database.GetConnectionURL(AURORA_USERNAME, AURORA_PASSWORD, AURORA_HOST, AURORA_DATABASE, AURORA_PORT))

	if err != nil {
		log.Println(err)

		return Response{Body: event.Name, StatusCode: 500}, nil
	}

	m.Up()

	fmt.Println("Received body: ", event.Name)

	return Response{Body: event.Name, StatusCode: 200}, nil
}

// main starts lambda using Handler
func main() {
	var err error

	keystore := adapters.NewKeyStore()

	AURORA_PASSWORD, err = keystore.ReadKey("EVS_AURORA_PASSWORD")

	if err != nil {
		log.Fatal(err)
	}

	lambda.Start(Handler)
}
