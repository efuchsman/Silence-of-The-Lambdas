package handlers

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	silence "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambdas"
	silenceofthelambdas "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambdas"
	"github.com/stretchr/testify/assert"
)

var (
	clientVictim = &silenceofthelambdas.Victim{
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

	clientVictim2 = &silenceofthelambdas.Victim{
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

	clientVictims = &silenceofthelambdas.Victims{
		Victims: []*silenceofthelambdas.Victim{
			clientVictim,
			clientVictim2,
		},
	}
)

func TestGetVictimsByKiller(t *testing.T) {
	testCases := []struct {
		description  string
		Killer       string
		client       *silence.TestClient
		expectedCode int
		expectedBody string
	}{
		{
			description: "Success: A Killer is returned",
			Killer:      "TestKiller",
			client: &silence.TestClient{
				ReturnVictimsByKillerData: clientVictims,
			},
			expectedCode: 200,
			expectedBody: `{"Victims":[{"killer":"TestKiller","full_name":"TestVictim","first_name":"Test","last_name":"Victim","actor":"Test Actor","movie":"Test Movie","cause_of_death":"Test","occupation":"Test","cannibalized":false,"image":"Test.jpg"},{"killer":"TestKiller","full_name":"TestVictim2","first_name":"Test","last_name":"Victim","actor":"Test Actor","movie":"Test Movie","cause_of_death":"Test","occupation":"Test","cannibalized":false,"image":"Test.jpg"}]}`,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			t.Log(tc.description)

			h := NewHandler(tc.client)

			tableName := "thisIsAMock"

			req := events.APIGatewayProxyRequest{}
			result := h.GetVictimsByKiller(req, tableName, tc.Killer)
			assert.NotNil(t, result)
			assert.Equal(t, tc.expectedCode, result.StatusCode)
			assert.Equal(t, tc.expectedBody, result.Body)
		})
	}
}
