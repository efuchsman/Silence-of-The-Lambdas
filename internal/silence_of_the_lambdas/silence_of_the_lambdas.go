package silenceofthelambdas

import (
	ddb "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambs_db"
)

type Client interface {
	ReturnKillerByFullName(fullName string, tableName string) (*Killer, error)
	ReturnVictimsByKiller(killerName string, tableName string) (*Victims, error)
}

type SilenceOfTheLambdasClient struct {
	db ddb.Client
}

func NewSilenceOfTheLambdasClient(db ddb.Client) *SilenceOfTheLambdasClient {
	return &SilenceOfTheLambdasClient{
		db: db,
	}
}
