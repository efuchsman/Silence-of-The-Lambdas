package main

import (
	"context"

	"os"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-lambda-go/events"
	"github.com/efuchsman/Silence-of-The-Lambdas/handlers"
	silence "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambdas"
	ddb "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambs_db"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")
	dynamoCreds := ddb.Credentials{
		AccessKeyID:     awsAccessKeyID,
		SecretAccessKey: awsSecretAccessKey,
	}

	db, err := ddb.NewSilenceOfTheLambsDB(awsRegion, awsRegion, &dynamoCreds)
	if err != nil {
		log.Fatal("Error creating SilenceOfTheLambsDB instance:", err)
	}

	sOTLClient := silence.NewSilenceOfTheLambdasClient(db)
	silenceHandler := handlers.NewHandler(sOTLClient)

	log.Printf("Received %s request for path: %s", request.HTTPMethod, request.Path)

	switch request.HTTPMethod {
	case "GET":
		response, err := handleGetRequest(request, silenceHandler, db)
		if err != nil {
			log.Printf("Error handling GET request: %v", err)
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Internal Server Error",
			}, nil
		}
		return *response, nil
	default:
		return events.APIGatewayProxyResponse{
			StatusCode: 405,
			Body:       "Method Not Allowed",
		}, nil
	}
}

func handleGetRequest(request events.APIGatewayProxyRequest, handler *handlers.Handler, db *ddb.SilenceOfTheLambsDB) (*events.APIGatewayProxyResponse, error) {
	switch request.Path {
	case "/":
		log.Println("Handling root path request.")
		return &events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "Hello, World!",
		}, nil
	case "/killers/full_name":
		log.Println("Handling /killers/full_name path request.")
		tableName := os.Getenv("DYNAMODB_TABLE_1_NAME")
		response := handler.GetKiller(request, tableName, db)
		return response, nil
	default:
		return &events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Not Found",
		}, nil
	}
}

func main() {
	lambda.Start(Handler)
}
