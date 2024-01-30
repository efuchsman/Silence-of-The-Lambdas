package dynamodb

type TestClient struct {
	ReturnKillerByFullNameData *Killer
	ReturnKillerByFullNameErr  error
}

func (c TestClient) ReturnKillerByFullName(fullName string, tableName string, db *SilenceOfTheLambsDB) (*Killer, error) {
	return c.ReturnKillerByFullNameData, c.ReturnKillerByFullNameErr
}
