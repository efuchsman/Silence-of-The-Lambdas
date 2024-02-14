package handlers

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	silence "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambdas"
	"github.com/stretchr/testify/assert"
)

var (
	killer = &silence.Killer{
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

func TestGetKiller(t *testing.T) {
	testCases := []struct {
		description  string
		FullName     string
		client       *silence.TestClient
		expectedCode int
		expectedBody string
	}{
		{
			description: "Success: A Killer is returned",
			FullName:    "TestKiller",
			client: &silence.TestClient{
				ReturnKillerByFullNameData: killer,
			},
			expectedCode: 200,
			expectedBody: `{"full_name":"TestKiller","first_name":"Test","last_name":"Killer","movie_actors":["Actor"],"movies":["Movie"],"nickname":"Test NickName","profession":"Test"}`,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			t.Log(tc.description)

			h := NewHandler(tc.client)

			tableName := "thisIsAMock"

			req := events.APIGatewayProxyRequest{}
			result := h.GetKiller(req, tableName, tc.FullName)
			assert.NotNil(t, result)
			assert.Equal(t, tc.expectedCode, result.StatusCode)
			assert.Equal(t, tc.expectedBody, result.Body)
		})
	}
}
