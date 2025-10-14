package mocks

import (
	"github.com/stretchr/testify/mock"
)

type TM3 struct {
	mock.Mock
}

func (m *TM3) GenerateRelatory(shouldSimulateError bool) ([]byte, error) {
	args := m.Called(shouldSimulateError)
	return args.Get(0).([]byte), args.Error(1)
}
