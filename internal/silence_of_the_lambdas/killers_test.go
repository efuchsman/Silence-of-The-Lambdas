package silenceofthelambdas

import (
	"fmt"
	"testing"

	silenceofthelambsdb "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambs_db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	killer = &silenceofthelambsdb.Killer{
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
		Image:      "Test.jpg",
	}
)

func TestReturnKillerByFullName(t *testing.T) {
	testCases := []struct {
		description string
		FullName    string
		client      *silenceofthelambsdb.TestClient
		expectedRes *Killer
		expectedErr error
	}{
		{
			description: "Success: A Killer is returned",
			FullName:    "TestKiller",
			client: &silenceofthelambsdb.TestClient{
				ReturnKillerByFullNameData: killer,
			},
			expectedRes: &Killer{
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
				Image:      "Test.jpg",
			},
			expectedErr: nil,
		},
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			t.Parallel()
			t.Log(tc.description)

			c := NewSilenceOfTheLambdasClient(tc.client)
			tableName := "DoesNotMatterForMock"

			result, err := c.ReturnKillerByFullName(tc.FullName, tableName)
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
