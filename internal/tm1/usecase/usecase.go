package usecase

import "github.com/charmingruby/spoc/internal/tm1/port"

type FetchInput struct {
	APIKey                      string
	ShouldSimulateAuthError     bool
	ShouldSimulateRelatoryError bool
}

type Service interface {
	Fetch(input FetchInput) ([]byte, error)
}

type UseCase struct {
	crm port.CRM
}

func New(
	crm port.CRM,
) *UseCase {
	return &UseCase{
		crm: crm,
	}
}
