package silenceofthelambsdb

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	killer = &Killer{
		FullName:  "TestKiller",
		FirstName: "Test",
		LastName:  "Killer",
		MovieActors: []string{
			"Actor",
		},
		Movies: []string{
			"Movie",
		},
		Nickname:   "Test NickName",
		Profession: "Test",
	}
)

func TestReturnKillerByFullName(t *testing.T) {
	testCases := []struct {
		description string
		FullName    string
		expectedRes *Killer
		client      *MockDBClient
		expectedErr error
	}{
		{
			description: "Success: DynamoDb returns a killer",
			FullName:    killer.FullName,
			client: &MockDBClient{
				GetItemOutput: &dynamodb.GetItemOutput{
					Item: map[string]*dynamodb.AttributeValue{
						"FullName": {
							S: aws.String("TestKiller"),
						},
						"FirstName": {
							S: aws.String("Test"),
						},
						"LastName": {
							S: aws.String("Killer"),
						},
						"MovieActors": {
							SS: []*string{
								aws.String("Actor"),
							},
						},
						"Movies": {
							SS: []*string{
								aws.String("Movie"),
							},
						},
						"Nickname": {
							S: aws.String("Test NickName"),
						},
						"Profession": {
							S: aws.String("Test"),
						},
					},
				},
			},
			expectedRes: killer,
			expectedErr: nil,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			t.Log(tc.description)

			db := &SilenceOfTheLambsDB{
				DynamoDB: tc.client,
			}

			result, err := db.ReturnKillerByFullName(tc.FullName, "")

			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NotNil(t, result)
				assert.NoError(t, err, tc.description)
				assert.Equal(t, tc.expectedRes, result)
				require.NoError(t, err)
			}
		})
	}
}
