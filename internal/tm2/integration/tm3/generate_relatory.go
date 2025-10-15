package tm3

import (
	"errors"

	"github.com/charmingruby/spoc/internal/tm1/data"
)

func (tt *TM3) GenerateRelatory(shouldSimulateError bool) ([]byte, error) {
	if shouldSimulateError {
		return nil, errors.New("unknown error")
	}

	return data.Relatory()
}
