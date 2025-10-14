package data

import (
	"encoding/json"

	"github.com/charmingruby/spoc/internal/tm1/model"
	"github.com/oklog/ulid/v2"
)

func ExternalRelatory() ([]byte, error) {
	r := model.ExternalRelatory{
		Page: 0,
	}

	for range 10 {
		r.Data = append(r.Data, model.ExternalRelatoryItem{
			ID:   ulid.Make().String(),
			Hash: ulid.Make().String(),
		})
	}

	return json.Marshal(r)
}
