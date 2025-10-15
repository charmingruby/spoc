package salesforce_test

import (
	"testing"

	"github.com/charmingruby/spoc/internal/tm1/integration/salesforce"
	"github.com/stretchr/testify/require"
)

func Test_Salesforce_GenerateRelatory(t *testing.T) {
	t.Run("should generate relatory sucessfully", func(t *testing.T) {
		sf := salesforce.New()

		validToken := "token"

		data, err := sf.GenerateRelatory(validToken, false)
		require.NoError(t, err)
		require.NotEmpty(t, data)
	})

	t.Run("should return ErrInvalidToken when token is invalid", func(t *testing.T) {
		sf := salesforce.New()

		invalidToken := "token"

		data, err := sf.GenerateRelatory(invalidToken, true)
		require.Error(t, err)
		require.ErrorIs(t, err, salesforce.ErrInvalidToken)
		require.Empty(t, data)
	})
}
