package usecase_test

import (
	"github.com/charmingruby/spoc/internal/tm1/usecase"
	"github.com/charmingruby/spoc/test/mocks"
)

type suite struct {
	sf      *mocks.Salesforce
	usecase *usecase.UseCase
}

func newSuite() *suite {
	sf := new(mocks.Salesforce)

	return &suite{
		sf:      sf,
		usecase: usecase.New(sf, usecase.Config{}),
	}
}
