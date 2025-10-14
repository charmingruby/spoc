package lambda_test

import (
	"testing"

	"github.com/charmingruby/spoc/internal/tm2/adapter/lambda"
	"github.com/stretchr/testify/require"
)

func Test_Lambda_GenerateRelatory(t *testing.T) {
	t.Run("should generate relatory sucessfully", func(t *testing.T) {
		l := lambda.New()

		data, err := l.GenerateRelatory(false)
		require.NoError(t, err)
		require.NotEmpty(t, data)
	})

	t.Run("should return error when there is any error", func(t *testing.T) {
		l := lambda.New()

		data, err := l.GenerateRelatory(true)
		require.Error(t, err)
		require.Empty(t, data)
	})
}
