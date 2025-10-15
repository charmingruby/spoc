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
		s := newSuite()

		apiKey := "apikey"
		token := "token"

		s.crm.On("Authenticate", apiKey, false).
			Return(token, nil)

		s.crm.On("GenerateRelatory", token, false).
			Return(data.Relatory())

		s.usecase.Config = usecase.Config{
			APIKey:                      apiKey,
			ShouldSimulateAuthError:     false,
			ShouldSimulateRelatoryError: false,
		}

		data, err := s.usecase.Fetch()

		require.NoError(t, err)
		require.NotEmpty(t, data)
	})

	t.Run("should return error when fails to authenticate", func(t *testing.T) {
		s := newSuite()

		apiKey := "apikey"

		s.crm.On("Authenticate", apiKey, true).
			Return("", errors.New("invalid api key"))

		s.usecase.Config = usecase.Config{
			APIKey:                  apiKey,
			ShouldSimulateAuthError: true,
		}

		data, err := s.usecase.Fetch()

		require.Error(t, err)
		require.Empty(t, data)
	})

	t.Run("should return error when fails to generate relatory", func(t *testing.T) {
		s := newSuite()

		apiKey := "apikey"
		token := "token"

		s.crm.On("Authenticate", apiKey, false).
			Return(token, nil)

		s.crm.On("GenerateRelatory", token, true).
			Return([]byte(nil), errors.New("invalid token"))

		s.usecase.Config = usecase.Config{
			APIKey:                      apiKey,
			ShouldSimulateAuthError:     false,
			ShouldSimulateRelatoryError: true,
		}

		data, err := s.usecase.Fetch()

		require.Error(t, err)
		require.Empty(t, data)
	})
}
