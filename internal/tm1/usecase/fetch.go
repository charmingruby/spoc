package usecase

func (u *UseCase) Fetch() ([]byte, error) {
	token, err := u.Salesforce.Authenticate(u.Config.APIKey, u.Config.ShouldSimulateAuthError)
	if err != nil {
		return nil, err
	}

	return u.Salesforce.GenerateRelatory(token, u.Config.ShouldSimulateRelatoryError)
}
