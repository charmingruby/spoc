package lambda

import (
	"errors"

	"github.com/charmingruby/spoc/internal/tm1/data"
)

func (l *Lambda) GenerateRelatory(shouldSimulateError bool) ([]byte, error) {
	if shouldSimulateError {
		return nil, errors.New("unknown error")
	}

	return data.Relatory()
}
