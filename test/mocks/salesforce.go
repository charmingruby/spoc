package mocks

import (
	"github.com/stretchr/testify/mock"
)

type Salesforce struct {
	mock.Mock
}

func (m *Salesforce) Authenticate(apiKey string, shouldSimulateError bool) (string, error) {
	args := m.Called(apiKey, shouldSimulateError)
	return args.String(0), args.Error(1)
}

func (m *Salesforce) GenerateRelatory(token string, shouldSimulateError bool) ([]byte, error) {
	args := m.Called(token, shouldSimulateError)
	return args.Get(0).([]byte), args.Error(1)
}
