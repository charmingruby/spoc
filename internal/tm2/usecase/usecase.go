package usecase

import "github.com/charmingruby/spoc/internal/tm2/port"

type FetchInput struct {
	ShouldSimulateRelatoryError bool
}

type Service interface {
	Fetch(input FetchInput) ([]byte, error)
}

type UseCase struct {
	tm3 port.TM3
}

func New(
	tm3 port.TM3,
) *UseCase {
	return &UseCase{
		tm3: tm3,
	}
}
