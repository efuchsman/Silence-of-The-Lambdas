package silenceofthelambsdb

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	victim = &Victim{
		Killer:       "TestKiller",
		FullName:     "TestVictim",
		FirstName:    "Test",
		LastName:     "Victim",
		Actor:        "Test Actor",
		Movie:        "Test Movie",
		CauseOfDeath: "Test",
		Occupation:   "Test",
		Cannibalized: false,
		Image:        "Test.jpg",
	}

	victim2 = &Victim{
		Killer:       "TestKiller",
		FullName:     "TestVictim2",
		FirstName:    "Test",
		LastName:     "Victim",
		Actor:        "Test Actor",
		Movie:        "Test Movie",
		CauseOfDeath: "Test",
		Occupation:   "Test",
		Cannibalized: false,
		Image:        "Test.jpg",
	}

	victims = &Victims{
		Victims: []*Victim{
			victim,
			victim2,
		},
	}
)

func TestReturnVictimsBYKiller(t *testing.T) {
	testCases := []struct {
		description string
		Killer      string
		client      *MockDBClient
		expectedRes *Victims
		expectedErr error
	}{
		{
			description: "Success: DynamoDB returns a killer's victims",
			Killer:      "TestKiller",
			client: &MockDBClient{
				QueryData: &dynamodb.QueryOutput{
					Items: []map[string]*dynamodb.AttributeValue{
						{
							"Killer": {
								S: aws.String(victim.Killer),
							},
							"FullName": {
								S: aws.String(victim.FullName),
							},
							"Actor": {
								S: aws.String(victim.Actor),
							},
							"Cannibalized": {
								BOOL: aws.Bool(victim.Cannibalized),
							},
							"CauseOfDeath": {
								S: aws.String(victim.CauseOfDeath),
							},
							"FirstName": {
								S: aws.String(victim.FirstName),
							},
							"Image": {
								S: aws.String(victim.Image),
							},
							"LastName": {
								S: aws.String(victim.LastName),
							},
							"Movie": {
								S: aws.String(victim.Movie),
							},
							"Occupation": {
								S: aws.String(victim.Occupation),
							},
						},
						{
							"Killer": {
								S: aws.String(victim2.Killer),
							},
							"FullName": {
								S: aws.String(victim2.FullName),
							},
							"Actor": {
								S: aws.String(victim2.Actor),
							},
							"Cannibalized": {
								BOOL: aws.Bool(victim2.Cannibalized),
							},
							"CauseOfDeath": {
								S: aws.String(victim2.CauseOfDeath),
							},
							"FirstName": {
								S: aws.String(victim2.FirstName),
							},
							"Image": {
								S: aws.String(victim2.Image),
							},
							"LastName": {
								S: aws.String(victim2.LastName),
							},
							"Movie": {
								S: aws.String(victim2.Movie),
							},
							"Occupation": {
								S: aws.String(victim2.Occupation),
							},
						},
					},
				},
			},
			expectedRes: victims,
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

			result, err := db.ReturnVictimsByKiller(tc.Killer, "")
			if tc.expectedErr != nil {
				assert.Equal(t, tc.expectedErr, err)
			} else {
				assert.NotNil(t, result)
				assert.NoError(t, err, tc.description)
				assert.True(t, reflect.DeepEqual(tc.expectedRes.Victims, result.Victims), tc.description)
				if !reflect.DeepEqual(tc.expectedRes.Victims, result.Victims) {
					fmt.Println("Expected:")
					for _, v := range tc.expectedRes.Victims {
						fmt.Printf("%+v\n", v)
					}

					fmt.Println("Actual  :")
					for _, v := range result.Victims {
						fmt.Printf("%+v\n", v)
					}
				}
				require.NoError(t, err)
			}
		})
	}
}
