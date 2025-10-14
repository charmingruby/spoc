package salesforce

import "errors"

var ErrInvalidAPIKey = errors.New("invalid api key")

func (s *Salesforce) Authenticate(apiKey string, shouldSimulateError bool) (string, error) {
	if shouldSimulateError {
		return "", ErrInvalidAPIKey
	}

	return "valid-token", nil
}
