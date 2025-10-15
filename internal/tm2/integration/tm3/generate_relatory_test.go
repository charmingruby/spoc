package tm3_test

import (
	"testing"

	"github.com/charmingruby/spoc/internal/tm2/integration/tm3"
	"github.com/stretchr/testify/require"
)

func Test_Lambda_GenerateRelatory(t *testing.T) {
	t.Run("should generate relatory sucessfully", func(t *testing.T) {
		tm := tm3.New()

		data, err := tm.GenerateRelatory(false)
		require.NoError(t, err)
		require.NotEmpty(t, data)
	})

	t.Run("should return error when there is any error", func(t *testing.T) {
		tm := tm3.New()

		data, err := tm.GenerateRelatory(true)
		require.Error(t, err)
		require.Empty(t, data)
	})
}
