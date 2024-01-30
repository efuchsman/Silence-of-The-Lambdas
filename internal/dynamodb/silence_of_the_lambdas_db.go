package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Client interface {
	ReturnKillerByFullName(fullName string, tableName string, db *SilenceOfTheLambsDB) (*Killer, error)
}

type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
}

type SilenceOfTheLambsDB struct {
	DynamoDB *dynamodb.DynamoDB
}

func NewSilenceOfTheLambsDB(region string, endpoint string, creds *Credentials) (*SilenceOfTheLambsDB, error) {
	awsCreds := credentials.NewStaticCredentials(creds.AccessKeyID, creds.SecretAccessKey, "")

	awsConfig := aws.Config{
		Region:      aws.String(region),
		Credentials: awsCreds,
	}

	if endpoint != "" {
		awsConfig.Endpoint = aws.String(endpoint)
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: awsCreds,
	})
	if err != nil {
		return nil, err
	}

	dynamoDBClient := dynamodb.New(sess)

	return &SilenceOfTheLambsDB{
		DynamoDB: dynamoDBClient,
	}, nil
}
