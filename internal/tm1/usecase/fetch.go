package usecase

func (u *UseCase) Fetch() ([]byte, error) {
	token, err := u.CRM.Authenticate(u.Config.APIKey, u.Config.ShouldSimulateAuthError)
	if err != nil {
		return nil, err
	}

	return u.CRM.GenerateRelatory(token, u.Config.ShouldSimulateRelatoryError)
}
