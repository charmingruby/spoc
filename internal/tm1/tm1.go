package tm1

import (
	"github.com/charmingruby/spoc/internal/shared/port"
	"github.com/charmingruby/spoc/internal/tm1/adapter/salesforce"
	"github.com/charmingruby/spoc/internal/tm1/usecase"
)

type Config = usecase.Config

func NewFetcher(cfg Config) port.Fetcher {
	sf := salesforce.New()
	return usecase.New(sf, cfg)
}
