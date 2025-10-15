package usecase

import "github.com/charmingruby/spoc/internal/shared/port"

type Service interface {
	CollectAndExport(fetchers []port.Fetcher) []error
}

type UseCase struct {
	// storage
}

func New() *UseCase {
	return &UseCase{}
}
