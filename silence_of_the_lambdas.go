package main

import (
	"context"
	"strings"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
	"github.com/efuchsman/Silence-of-The-Lambdas/handlers"
	"github.com/joho/godotenv"
	
	silence "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambdas"
	ddb "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambs_db"
	log "github.com/sirupsen/logrus"
)

// main should be the core function and initialize critical dependencies, process flags, init logging etc. 
// Additionally the file should be called `main.go`
func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)





	lambda.Start(Handler)
	log.StandardLogger().Exit(0) // Odd to see a hard exit for a logger
}

// You're passing this function to lamda.Start. It should be initialized here if it's a critical dependency.

type LambdaHandler struct {
	db *ddb.SilenceOfTheLambsDB
	// TODO
}

// func (h *LambdaHandler) handleGetRequest(...) {...}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// A lot of this should be moved into main, especially when calling log.Fail.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// These should all be handled in main()
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

	// This can be passed directly to the handling function. It's hidden where it currently is.
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

// More reason to separate the above function - lots of unecessaryly hidden functionality and overall flow patterns. It looks like the above should be a new handler struct, which is passed into the lambda.
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
