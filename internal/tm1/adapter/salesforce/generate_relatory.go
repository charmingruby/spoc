package salesforce

import (
	"errors"

	"github.com/charmingruby/spoc/internal/tm1/data"
)

var ErrInvalidToken = errors.New("expired token")

func (s *Salesforce) GenerateRelatory(token string, shouldSimulateError bool) ([]byte, error) {
	if shouldSimulateError {
		return nil, ErrInvalidToken
	}

	return data.Relatory()
}
