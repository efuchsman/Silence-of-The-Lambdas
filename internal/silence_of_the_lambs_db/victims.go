package silenceofthelambsdb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	log "github.com/sirupsen/logrus"
)

type Victim struct {
	FullName     string `json:"full_name" dynamodbav:"FullName"`
	FirstName    string `json:"first_name" dynamodbav:"FirstName"`
	LastName     string `json:"last_name" dynamodbav:"LastName"`
	Actor        string `json:"actor" dynamodbav:"Actor"`
	Movie        string `json:"movie" dynamodbav:"Movie"`
	CauseOfDeath string `json:"cause_of_death" dynamodbav:"CauseOfDeath"`
	Occupation   string `json:"occupation" dynamodbav:"Occupation"`
	Cannibalized bool   `json:"cannibalized" dynamodbav:"Cannibalized"`
}

type Victims struct {
	Victims []*Victim
}

// ReturnVictimsByKiller takes in a Killer fullName and table input with no spaces and calls on Dynamodb to return the victims items
func (ddb *SilenceOfTheLambsDB) ReturnVictimsByKiller(killerName string, tableName string) (*Victims, error) {
	input := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("Killer = :killerName"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":killerName": {
				S: aws.String(killerName),
			},
		},
	}

	result, err := ddb.DynamoDB.Query(input)
	if err != nil {
		log.Println("Error getting item:", err)
		return nil, err
	}

	victims := &Victims{Victims: make([]*Victim, 0)}

	for _, item := range result.Items {
		victim := &Victim{}
		if err = dynamodbattribute.UnmarshalMap(item, victim); err != nil {
			log.Println("Error unmarshaling item:", err)
			return nil, err
		}
		victims.Victims = append(victims.Victims, victim)
	}
	return victims, nil
}
