package mocks

import (
	"github.com/stretchr/testify/mock"
)

type CRM struct {
	mock.Mock
}

func (m *CRM) Authenticate(apiKey string, shouldSimulateError bool) (string, error) {
	args := m.Called(apiKey, shouldSimulateError)
	return args.String(0), args.Error(1)
}

func (m *CRM) GenerateRelatory(token string, shouldSimulateError bool) ([]byte, error) {
	args := m.Called(token, shouldSimulateError)
	return args.Get(0).([]byte), args.Error(1)
}
