package collector

import (
	"github.com/charmingruby/spoc/internal/collector/usecase"
	"github.com/charmingruby/spoc/internal/shared/fetcher"
	"github.com/charmingruby/spoc/internal/shared/storage"
)

type Config = usecase.Config

type Collector struct {
	usecase usecase.Service
}

func New(cfg Config, storage storage.Storage) *Collector {
	return &Collector{
		usecase: usecase.New(cfg, storage),
	}
}

func (c *Collector) Run(fetchers []fetcher.Fetcher) []error {
	return c.usecase.CollectAndExport(fetchers)
}
