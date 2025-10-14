package tm2

import (
	"github.com/charmingruby/spoc/internal/shared/port"
	"github.com/charmingruby/spoc/internal/tm1/adapter/salesforce"
	"github.com/charmingruby/spoc/internal/tm1/usecase"
)

func NewFetcher() port.Fetcher[usecase.FetchInput] {
	sf := salesforce.New()
	return usecase.New(sf)
}
