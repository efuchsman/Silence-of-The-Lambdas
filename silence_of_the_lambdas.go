package main

import (
	"fmt"
	"log"
	"os"

	ddb "github.com/efuchsman/Silence-of-The-Lambdas/internal/dynamodb"
	"github.com/joho/godotenv"
)

func main() {
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

	db, err := ddb.NewSilenceOfTheLambsDB(awsRegion, "http://localhost:8000", &dynamoCreds)
	if err != nil {
		log.Fatal("Error creating SilenceOfTheLambsDB instance:", err)
	}

	fmt.Printf("Database created: %+v", db)
}
