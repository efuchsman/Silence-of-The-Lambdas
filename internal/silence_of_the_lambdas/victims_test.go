package silenceofthelambdas

import (
	"fmt"
	"testing"

	silenceofthelambsdb "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambs_db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	dbVictim = &silenceofthelambsdb.Victim{
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

	dbVictim2 = &silenceofthelambsdb.Victim{
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

	dbVictims = &silenceofthelambsdb.Victims{
		Victims: []*silenceofthelambsdb.Victim{
			dbVictim,
			dbVictim2,
		},
	}

	clientVictim = &Victim{
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

	clientVictim2 = &Victim{
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

	clientVictims = &Victims{
		Victims: []*Victim{
			clientVictim,
			clientVictim2,
		},
	}
)

func TestReturnVictimsByKiller(t *testing.T) {
	testCases := []struct {
		description string
		Killer      string
		client      *silenceofthelambsdb.TestClient
		expectedRes *Victims
		expectedErr error
	}{
		{
			description: "Success: A Killer's victims are returned",
			Killer:      "TestKiller",
			client: &silenceofthelambsdb.TestClient{
				ReturnVictimsByKillerData: dbVictims,
			},
			expectedRes: clientVictims,
			expectedErr: nil,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			t.Log(tc.description)

			c := NewSilenceOfTheLambdasClient(tc.client)
			tableName := "DoesNotMatterForMock"

			result, err := c.ReturnVictimsByKiller(tc.Killer, tableName)
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
