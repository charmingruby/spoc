package collector

import (
	"github.com/charmingruby/spoc/internal/collector/usecase"
	"github.com/charmingruby/spoc/internal/shared/port"
)

type Collector struct {
	usecase usecase.Service
}

func New() *Collector {
	return &Collector{
		usecase: usecase.New(),
	}
}

func (c *Collector) Run(fetchers []port.Fetcher) []error {
	return c.usecase.CollectAndExport(fetchers)
}
