package usecase

func (u *UseCase) Fetch(in FetchInput) ([]byte, error) {
	return u.tm3.GenerateRelatory(in.ShouldSimulateRelatoryError)
}
