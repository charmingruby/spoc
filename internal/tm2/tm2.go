package tm2

import (
	"github.com/charmingruby/spoc/internal/shared/port"
	"github.com/charmingruby/spoc/internal/tm2/adapter/lambda"
	"github.com/charmingruby/spoc/internal/tm2/usecase"
)

type Config = usecase.Config

func NewFetcher(cfg Config) port.Fetcher {
	lambda := lambda.New()
	return usecase.New(lambda, cfg)
}
