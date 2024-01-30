package dynamodb

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Killer struct {
	FullName    string   `json:"full_name"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	MovieActors []string `json:"movie_actors"`
	Movies      []string `json:"movies"`
	Nickname    string   `json:"nickname"`
	Profession  string   `json:"profession"`
}

func ReturnKillerByFullName(fullName string, tableName string, db *SilenceOfTheLambsDB) (*Killer, error) {
	// Prepare input for GetItem
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"FullName": {
				S: aws.String(fullName),
			},
		},
	}

	// GetItem request using the provided DynamoDB client
	result, err := db.DynamoDB.GetItem(input)
	if err != nil {
		log.Fatal("Error getting item:", err)
		return nil, err
	}

	// Parse the result into Killers struct
	item := &Killer{}
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		log.Fatal("Error unmarshalling item:", err)
		return nil, err
	}

	return item, nil
}
