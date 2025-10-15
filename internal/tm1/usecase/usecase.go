package usecase

import "github.com/charmingruby/spoc/internal/tm1/integration"

type Config struct {
	APIKey                      string
	ShouldSimulateAuthError     bool
	ShouldSimulateRelatoryError bool
}

type Service interface {
	Fetch() ([]byte, error)
}

type UseCase struct {
	Salesforce integration.Salesforce
	Config     Config
}

func New(
	sf integration.Salesforce,
	cfg Config,
) *UseCase {
	return &UseCase{
		Salesforce: sf,
		Config:     cfg,
	}
}
