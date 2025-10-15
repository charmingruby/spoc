package usecase

import (
	"context"
	"sync"

	"github.com/charmingruby/spoc/internal/shared/fetcher"
)

type fetchResult struct {
	data []byte
	err  error
}

func (u *UseCase) CollectAndExport(fetchers []fetcher.Fetcher) []error {
	ctx, cancel := context.WithTimeout(context.Background(), u.Config.Timeout)
	defer cancel()

	var (
		wg      sync.WaitGroup
		mu      sync.Mutex
		errors  []error
		results = make(chan fetchResult, len(fetchers))
		uploads = make(chan []byte, len(fetchers))
	)

	if err := u.Storage.Reset(ctx); err != nil {
		return []error{err}
	}

	numWorkers := u.Config.MaxWorkers
	if numWorkers == 0 {
		numWorkers = len(fetchers)
	}

	jobs := make(chan fetcher.Fetcher, len(fetchers))
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for fetcher := range jobs {
				done := make(chan fetchResult, 1)

				go func() {
					data, err := fetcher.Fetch()
					done <- fetchResult{data: data, err: err}
				}()

				select {
				case result := <-done:
					results <- result
				case <-ctx.Done():
					results <- fetchResult{err: ctx.Err()}
					return
				}
			}
		}()
	}

	var uploadWg sync.WaitGroup
	maxUploadWorkers := 10
	for range maxUploadWorkers {
		uploadWg.Add(1)
		go func() {
			defer uploadWg.Done()
			for data := range uploads {
				if err := u.Storage.Upload(ctx, data); err != nil {
					mu.Lock()
					errors = append(errors, err)
					mu.Unlock()
				}
			}
		}()
	}

	go func() {
		defer close(jobs)
		for _, f := range fetchers {
			select {
			case jobs <- f:
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		for result := range results {
			if result.err != nil {
				mu.Lock()
				errors = append(errors, result.err)
				mu.Unlock()
				continue
			}

			uploads <- result.data
		}
		close(uploads)
	}()

	wg.Wait()
	close(results)

	uploadWg.Wait()

	if len(errors) > 0 {
		return errors
	}

	return nil
}
