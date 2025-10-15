package storage

import "context"

type Storage interface {
	Upload(ctx context.Context, data []byte) error
	Reset(ctx context.Context) error
}
