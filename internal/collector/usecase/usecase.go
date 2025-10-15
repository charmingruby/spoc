package usecase

import (
	"time"

	"github.com/charmingruby/spoc/internal/shared/fetcher"
	"github.com/charmingruby/spoc/internal/shared/storage"
)

type Config struct {
	Timeout    time.Duration
	MaxWorkers int
}

type Service interface {
	CollectAndExport(fetchers []fetcher.Fetcher) []error
}

type UseCase struct {
	Storage storage.Storage
	Config  Config
}

func New(cfg Config, storage storage.Storage) *UseCase {
	return &UseCase{
		Config:  cfg,
		Storage: storage,
	}
}
