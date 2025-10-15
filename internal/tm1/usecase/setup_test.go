package usecase_test

import (
	"github.com/charmingruby/spoc/internal/tm1/usecase"
	"github.com/charmingruby/spoc/test/mocks"
)

type suite struct {
	crm     *mocks.CRM
	usecase *usecase.UseCase
}

func newSuite() *suite {
	crm := new(mocks.CRM)

	return &suite{
		crm:     crm,
		usecase: usecase.New(crm, usecase.Config{}),
	}
}
