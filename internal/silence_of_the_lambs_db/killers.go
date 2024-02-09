package silenceofthelambsdb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	log "github.com/sirupsen/logrus"
)

type Killer struct {
	FullName    string   `json:"full_name" dynamodbav:"FullName"`
	FirstName   string   `json:"first_name" dynamodbav:"FirstName"`
	LastName    string   `json:"last_name" dynamodbav:"LastName"`
	MovieActors []string `json:"movie_actors" dynamodbav:"MovieActors"`
	Movies      []string `json:"movies" dynamodbav:"Movies"`
	Nickname    string   `json:"nickname" dynamodbav:"Nickname"`
	Profession  string   `json:"profession" dynamodbav:"Profession"`
}

// ReturnKillerByFullName takes in a fullName and table input with no spaces and calls on Dynamodb to return the item
func (ddb *SilenceOfTheLambsDB) ReturnKillerByFullName(fullName string, tableName string) (*Killer, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"FullName": {
				S: aws.String(fullName),
			},
		},
	}

	result, err := ddb.DynamoDB.GetItem(input)
	if err != nil {
		log.Println("Error getting item:", err)
		return nil, err
	}

	item := &Killer{}
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		log.Println("Error unmarshalling item:", err)
		return nil, err
	}

	return item, nil
}
