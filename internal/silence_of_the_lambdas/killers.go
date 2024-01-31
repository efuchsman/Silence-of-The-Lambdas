package silenceofthelambdas

import "github.com/efuchsman/Silence-of-The-Lambdas/internal/dynamodb"

type Killer struct {
	FullName    string   `json:"full_name"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	MovieActors []string `json:"movie_actors"`
	Movies      []string `json:"movies"`
	Nickname    string   `json:"nickname"`
	Profession  string   `json:"profession"`
}

func (c *SilenceOfTheLambdasClient) ReturnKillerByFullName(fullName string, tableName string, db *dynamodb.SilenceOfTheLambsDB) (*Killer, error) {
	return nil, nil
}
