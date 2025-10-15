package usecase

func (u *UseCase) Fetch() ([]byte, error) {
	return u.TM3.GenerateRelatory(u.Config.ShouldSimulateRelatoryError)
}
