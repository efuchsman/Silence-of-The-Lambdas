package main

import (
	"context"
	"strings"

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
	log.SetOutput(os.Stdout)
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

	db, err := ddb.NewSilenceOfTheLambsDB(awsRegion, "", &dynamoCreds)
	if err != nil {
		log.Fatal("Error creating SilenceOfTheLambsDB instance:", err)
	}

	sOTLClient := silence.NewSilenceOfTheLambdasClient(db)
	silenceHandler := handlers.NewHandler(sOTLClient)

	log.WithFields(log.Fields{
		"headers":    request.Headers,
		"path":       request.Path,
		"httpMethod": request.HTTPMethod,
	}).Info("Request details")

	if request.HTTPMethod == "GET" {
		response, err := handleGetRequest(request, silenceHandler, db)
		if err != nil {
			log.WithFields(log.Fields{
				"error":   err,
				"request": request,
			}).Error("Error handling GET request")
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Internal Server Error",
			}, nil
		}
		return *response, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 405,
		Body:       "Method Not Allowed",
	}, nil
}

func handleGetRequest(request events.APIGatewayProxyRequest, handler *handlers.Handler, db *ddb.SilenceOfTheLambsDB) (*events.APIGatewayProxyResponse, error) {
	cleanedPath := strings.TrimSpace(request.Path)
	log.WithFields(log.Fields{
		"cleanedPath": cleanedPath,
		"headers":     request.Headers,
		"path":        request.Path,
		"httpMethod":  request.HTTPMethod,
	}).Info("Request details")

	switch cleanedPath {
	case "/":
		log.WithFields(log.Fields{
			"path":    request.Path,
			"request": request,
		}).Info("Handling root path request")
		return &events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "Welcome to Silence of The Lambs API!",
		}, nil
	case "/killers/full_name":
		tableName := os.Getenv("DYNAMODB_TABLE_1_NAME")
		response := handler.GetKiller(request, tableName, db)
		log.WithFields(log.Fields{
			"path":    request.Path,
			"request": request,
		}).Info("Handling /killers/full_name request")
		return response, nil
	default:
		log.WithFields(log.Fields{
			"cleanedPath": cleanedPath,
			"request":     request,
		}).Warn("Unknown path requested")
		return &events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Not Found",
		}, nil
	}
}

func main() {
	lambda.Start(Handler)
	log.StandardLogger().Exit(0)
}
