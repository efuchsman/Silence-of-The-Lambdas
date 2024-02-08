package silenceofthelambsdb

import "github.com/aws/aws-sdk-go/service/dynamodb"

type MockDBClient struct {
	GetItemOutput *dynamodb.GetItemOutput
	GetItemError  error
}

// Used to mock a DynamoDB request
func (m *MockDBClient) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return m.GetItemOutput, m.GetItemError
}
