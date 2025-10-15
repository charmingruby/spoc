package tm1

import (
	"github.com/charmingruby/spoc/internal/shared/fetcher"
	"github.com/charmingruby/spoc/internal/tm1/integration/salesforce"
	"github.com/charmingruby/spoc/internal/tm1/usecase"
)

type Config = usecase.Config

func NewFetcher(cfg Config) fetcher.Fetcher {
	sf := salesforce.New()
	return usecase.New(sf, cfg)
}
