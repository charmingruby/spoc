package usecase

func (u *UseCase) Fetch(in FetchInput) ([]byte, error) {
	token, err := u.crm.Authenticate(in.APIKey, in.ShouldSimulateError)
	if err != nil {
		return nil, err
	}

	return u.crm.GenerateRelatory(token, in.ShouldSimulateError)
}
