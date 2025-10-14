package usecase_test

import (
	"errors"
	"testing"

	"github.com/charmingruby/spoc/internal/tm1/data"
	"github.com/charmingruby/spoc/internal/tm1/usecase"
	"github.com/stretchr/testify/require"
)

func Test_UseCase_Fetch(t *testing.T) {
	t.Run("should fetch data sucessfully", func(t *testing.T) {
		m := newSuite()

		apiKey := "apikey"
		token := "token"

		m.crm.On("Authenticate", apiKey, false).
			Return(token, nil)

		m.crm.On("GenerateRelatory", token, false).
			Return(data.Relatory())

		data, err := m.usecase.Fetch(usecase.FetchInput{
			APIKey:                      apiKey,
			ShouldSimulateAuthError:     false,
			ShouldSimulateRelatoryError: false,
		})

		require.NoError(t, err)
		require.NotEmpty(t, data)
	})

	t.Run("should return error when fails to authenticate", func(t *testing.T) {
		m := newSuite()

		apiKey := "apikey"

		m.crm.On("Authenticate", apiKey, true).
			Return("", errors.New("invalid api key"))

		data, err := m.usecase.Fetch(usecase.FetchInput{
			APIKey:                  apiKey,
			ShouldSimulateAuthError: true,
		})

		require.Error(t, err)
		require.Empty(t, data)
	})

	t.Run("should return error when fails to generate relatory", func(t *testing.T) {
		m := newSuite()

		apiKey := "apikey"
		token := "token"

		m.crm.On("Authenticate", apiKey, false).
			Return(token, nil)

		m.crm.On("GenerateRelatory", token, true).
			Return([]byte(nil), errors.New("invalid token"))

		data, err := m.usecase.Fetch(usecase.FetchInput{
			APIKey:                      apiKey,
			ShouldSimulateAuthError:     false,
			ShouldSimulateRelatoryError: true,
		})

		require.Error(t, err)
		require.Empty(t, data)
	})
}
