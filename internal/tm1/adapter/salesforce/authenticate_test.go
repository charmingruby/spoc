package salesforce_test

import (
	"testing"

	"github.com/charmingruby/spoc/internal/tm1/adapter/salesforce"
	"github.com/stretchr/testify/require"
)

func Test_Salesforce_Authenticate(t *testing.T) {
	t.Run("should authenticate sucessfully", func(t *testing.T) {
		sf := salesforce.New()

		validAPIKey := "APIKey"

		data, err := sf.Authenticate(validAPIKey, false)
		require.NoError(t, err)
		require.NotEmpty(t, data)
	})

	t.Run("should return ErrInvalidAPIKey when api key is invalid", func(t *testing.T) {
		sf := salesforce.New()

		invalidAPIKey := "APIKey"

		data, err := sf.Authenticate(invalidAPIKey, true)
		require.Error(t, err)
		require.ErrorIs(t, err, salesforce.ErrInvalidAPIKey)
		require.Empty(t, data)
	})
}
