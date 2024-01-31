package silenceofthelambdas

import ddb "github.com/efuchsman/Silence-of-The-Lambdas/internal/silence_of_the_lambs_db"

type TestClient struct {
	ReturnKillerByFullNameData *Killer
	ReturnKillerByFullNameErr  error
}

func (c TestClient) ReturnKillerByFullName(fullName string, tableName string, db *ddb.SilenceOfTheLambsDB) (*Killer, error) {
	return c.ReturnKillerByFullNameData, c.ReturnKillerByFullNameErr
}
