package data

import (
	"encoding/json"

	"github.com/charmingruby/spoc/internal/tm1/model"
	"github.com/oklog/ulid/v2"
)

func Relatory() ([]byte, error) {
	r := model.Relatory{
		Page: 0,
	}

	for range 10 {
		r.Data = append(r.Data, model.RelatoryItem{
			ID:   ulid.Make().String(),
			Hash: ulid.Make().String(),
		})
	}

	return json.Marshal(r)
}
