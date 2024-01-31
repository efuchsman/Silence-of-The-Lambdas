package silenceofthelambdas

import (
	ddb "github.com/efuchsman/Silence-of-The-Lambdas/internal/dynamodb"
)

type Client interface {
}

type SilenceOfTheLambdasClient struct {
	db ddb.Client
}

func NewSilenceOfTheLambdasClient(db ddb.Client) *SilenceOfTheLambdasClient {
	return &SilenceOfTheLambdasClient{
		db: db,
	}
}
