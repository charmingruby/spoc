package usecase_test

import (
	"errors"
	"testing"

	"github.com/charmingruby/spoc/internal/tm2/data"
	"github.com/charmingruby/spoc/internal/tm2/usecase"
	"github.com/stretchr/testify/require"
)

func Test_UseCase_Fetch(t *testing.T) {
	t.Run("should fetch data sucessfully", func(t *testing.T) {
		s := newSuite()

		s.tm3.On("GenerateRelatory", false).
			Return(data.Relatory())

		s.usecase.Config = usecase.Config{
			ShouldSimulateRelatoryError: false,
		}

		data, err := s.usecase.Fetch()

		require.NoError(t, err)
		require.NotEmpty(t, data)
	})

	t.Run("should return error when fails to generate relatory", func(t *testing.T) {
		s := newSuite()

		s.tm3.On("GenerateRelatory", true).
			Return([]byte(nil), errors.New("unknown error"))

		s.usecase.Config = usecase.Config{
			ShouldSimulateRelatoryError: true,
		}

		data, err := s.usecase.Fetch()

		require.Error(t, err)
		require.Empty(t, data)
	})
}
