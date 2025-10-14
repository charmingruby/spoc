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
		m := newSuite()

		m.tm3.On("GenerateRelatory", false).
			Return(data.Relatory())

		data, err := m.usecase.Fetch(usecase.FetchInput{
			ShouldSimulateRelatoryError: false,
		})

		require.NoError(t, err)
		require.NotEmpty(t, data)
	})

	t.Run("should return error when fails to generate relatory", func(t *testing.T) {
		m := newSuite()

		m.tm3.On("GenerateRelatory", true).
			Return([]byte(nil), errors.New("unknown error"))

		data, err := m.usecase.Fetch(usecase.FetchInput{
			ShouldSimulateRelatoryError: true,
		})

		require.Error(t, err)
		require.Empty(t, data)
	})
}
