package mocks

import "github.com/stretchr/testify/mock"

type Fetcher struct {
	mock.Mock
}

func (m *Fetcher) Fetch() ([]byte, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]byte), args.Error(1)
}
