package model_test

import (
	"encoding/json"
	"testing"

	"github.com/charmingruby/spoc/internal/tm1/model"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/require"
)

func Test_ExternalRelatory(t *testing.T) {
	t.Run("should marshal sucessfully", func(t *testing.T) {
		relatory := model.Relatory{
			Page: 0,
			Data: []model.RelatoryItem{
				{
					ID:   ulid.Make().String(),
					Hash: ulid.Make().String(),
				},
			},
		}

		data, err := json.Marshal(relatory)
		require.NoError(t, err)

		var parsedRelatory model.Relatory
		err = json.Unmarshal(data, &parsedRelatory)
		require.NoError(t, err)
		require.Equal(t, parsedRelatory, relatory)
	})
}
