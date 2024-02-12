package silenceofthelambsdb

type TestClient struct {
	ReturnKillerByFullNameData *Killer
	ReturnKillerByFullNameErr  error

	ReturnVictimsByKillerData *Victims
	ReturnVictimsByKillerErr  error
}

func (c *TestClient) ReturnKillerByFullName(fullName string, tableName string) (*Killer, error) {
	return c.ReturnKillerByFullNameData, c.ReturnKillerByFullNameErr
}

func (c *TestClient) ReturnVictimsByKiller(killerName string, tableName string) (*Victims, error) {
	return c.ReturnVictimsByKillerData, c.ReturnVictimsByKillerErr
}
