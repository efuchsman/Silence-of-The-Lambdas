package silenceofthelambsdb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Client interface {
	ReturnKillerByFullName(fullName string, tableName string) (*Killer, error)
}

type DBClient interface {
	GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
}

type SilenceOfTheLambsDB struct {
	DynamoDB DBClient
}

func NewSilenceOfTheLambsDB(region string, endpoint string) (*SilenceOfTheLambsDB, error) {
	awsConfig := aws.Config{
		Region: aws.String(region),
	}

	if endpoint != "" {
		awsConfig.Endpoint = aws.String(endpoint)
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	dynamoDBClient := dynamodb.New(sess)

	return &SilenceOfTheLambsDB{
		DynamoDB: dynamoDBClient,
	}, nil
}
