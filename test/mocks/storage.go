package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type Storage struct {
	mock.Mock
}

func (m *Storage) Reset(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *Storage) Upload(ctx context.Context, data []byte) error {
	args := m.Called(ctx, data)
	return args.Error(0)
}
