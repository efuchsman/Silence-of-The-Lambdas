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

type LambdaHandler struct {
	db *ddb.SilenceOfTheLambsDB
	// TODO
}

func (h *LambdaHandler) Handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {

	log.WithFields(log.Fields{
		"headers":    request.Headers,
		"path":       request.Path,
		"httpMethod": request.HTTPMethod,
	}).Info("Request details")

	switch request.HTTPMethod {
	case "GET":
		return h.handleGetRequest(request, h.db)
	default:
		return &events.APIGatewayProxyResponse{
			StatusCode: 405,
			Body:       "Method Not Allowed",
		}, nil
	}
}

func (h *LambdaHandler) handleGetRequest(request events.APIGatewayProxyRequest, db *ddb.SilenceOfTheLambsDB) (*events.APIGatewayProxyResponse, error) {
	log.WithFields(log.Fields{
		"headers":    request.Headers,
		"path":       request.Path,
		"params":     request.PathParameters,
		"resource":   request.Resource,
		"httpMethod": request.HTTPMethod,
	}).Info("Request details")

	switch request.Resource {
	case "/":
		log.WithFields(log.Fields{
			"path":    request.Path,
			"request": request,
		}).Info("Handling root path request")
		return &events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       "Welcome to Silence of The Lambdas API!",
		}, nil
	case "/killers/{full_name}":
		tableName := os.Getenv("DYNAMODB_TABLE_1_NAME")
		handler := handlers.NewHandler(silence.NewSilenceOfTheLambdasClient(h.db))

		fullName, ok := request.PathParameters["full_name"]
		if !ok {
			log.Warn("Missing path parameter 'full_name'")
			return &events.APIGatewayProxyResponse{
				StatusCode: 400,
				Body:       "Bad Request: Missing 'full_name' parameter",
			}, nil
		}

		log.WithFields(log.Fields{
			"path":       request.Path,
			"params":     request.PathParameters,
			"request":    request,
			"table_name": tableName,
		}).Info("Handling /killers/full_name request")

		response := handler.GetKiller(request, tableName, fullName)

		return response, nil
	default:
		log.WithFields(log.Fields{
			"path":    request.Path,
			"params":  request.PathParameters,
			"request": request,
		}).Warn("Unknown path requested")
		return &events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "Not Found",
		}, nil
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	awsRegion := os.Getenv("AWS_REGION")
	db, err := ddb.NewSilenceOfTheLambsDB(awsRegion, "")
	if err != nil {
		log.Fatal("Error creating SilenceOfTheLambsDB instance:", err)
	}

	lamHandler := &LambdaHandler{
		db: db,
	}

	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	lambda.Start(lamHandler.Handler)
	log.StandardLogger().Exit(0)
}
