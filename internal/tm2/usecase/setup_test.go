package usecase_test

import (
	"github.com/charmingruby/spoc/internal/tm2/usecase"
	"github.com/charmingruby/spoc/test/mocks"
)

type suite struct {
	tm3     *mocks.TM3
	usecase *usecase.UseCase
}

func newSuite() *suite {
	tm3 := new(mocks.TM3)

	return &suite{
		tm3:     tm3,
		usecase: usecase.New(tm3, usecase.Config{}),
	}
}
