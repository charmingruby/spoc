package tm2

import (
	"github.com/charmingruby/spoc/internal/shared/fetcher"
	"github.com/charmingruby/spoc/internal/tm2/integration/tm3"
	"github.com/charmingruby/spoc/internal/tm2/usecase"
)

type Config = usecase.Config

func NewFetcher(cfg Config) fetcher.Fetcher {
	tm3 := tm3.New()
	return usecase.New(tm3, cfg)
}
