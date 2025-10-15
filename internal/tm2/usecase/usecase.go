package usecase

import "github.com/charmingruby/spoc/internal/tm2/integration"

type Config struct {
	ShouldSimulateRelatoryError bool
}

type Service interface {
	Fetch() ([]byte, error)
}

type UseCase struct {
	TM3    integration.TM3
	Config Config
}

func New(
	tm3 integration.TM3,
	cfg Config,
) *UseCase {
	return &UseCase{
		TM3:    tm3,
		Config: cfg,
	}
}
