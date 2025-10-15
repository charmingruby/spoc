package usecase

import "github.com/charmingruby/spoc/internal/shared/port"

func (u *UseCase) CollectAndExport(fetchers []port.Fetcher) []error {
	var errors []error

	for _, f := range fetchers {
		data, err := f.Fetch()
		if err != nil {
			errors = append(errors, err)
			continue
		}

		println(string(data))

		// export data
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}
