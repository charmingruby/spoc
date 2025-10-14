package tm2

import (
	"github.com/charmingruby/spoc/internal/shared/port"
	"github.com/charmingruby/spoc/internal/tm2/adapter/lambda"
	"github.com/charmingruby/spoc/internal/tm2/usecase"
)

func NewFetcher() port.Fetcher[usecase.FetchInput] {
	lambda := lambda.New()
	return usecase.New(lambda)
}
