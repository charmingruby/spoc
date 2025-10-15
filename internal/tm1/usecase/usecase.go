package usecase

import "github.com/charmingruby/spoc/internal/tm1/port"

type Config struct {
	APIKey                      string
	ShouldSimulateAuthError     bool
	ShouldSimulateRelatoryError bool
}

type Service interface {
	Fetch() ([]byte, error)
}

type UseCase struct {
	CRM    port.CRM
	Config Config
}

func New(
	CRM port.CRM,
	Config Config,
) *UseCase {
	return &UseCase{
		CRM:    CRM,
		Config: Config,
	}
}
