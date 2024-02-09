package silenceofthelambsdb

type TestClient struct {
	ReturnKillerByFullNameData *Killer
	ReturnKillerByFullNameErr  error
}

func (c *TestClient) ReturnKillerByFullName(fullName string, tableName string) (*Killer, error) {
	return c.ReturnKillerByFullNameData, c.ReturnKillerByFullNameErr
}
