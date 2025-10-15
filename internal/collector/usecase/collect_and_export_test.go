package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/charmingruby/spoc/internal/collector/usecase"
	"github.com/charmingruby/spoc/internal/shared/fetcher"
	"github.com/charmingruby/spoc/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCollectAndExport_Success(t *testing.T) {
	storage := new(mocks.Storage)
	fetcher1 := new(mocks.Fetcher)
	fetcher2 := new(mocks.Fetcher)

	data1 := []byte(`{"id":1}`)
	data2 := []byte(`{"id":2}`)

	storage.On("Reset", mock.Anything).Return(nil)
	fetcher1.On("Fetch").Return(data1, nil)
	fetcher2.On("Fetch").Return(data2, nil)
	storage.On("Upload", mock.Anything, data1).Return(nil)
	storage.On("Upload", mock.Anything, data2).Return(nil)

	uc := &usecase.UseCase{
		Config: usecase.Config{
			Timeout:    5 * time.Second,
			MaxWorkers: 2,
		},
		Storage: storage,
	}

	fetchers := []fetcher.Fetcher{fetcher1, fetcher2}
	errs := uc.CollectAndExport(fetchers)

	assert.Nil(t, errs)
	storage.AssertExpectations(t)
	fetcher1.AssertExpectations(t)
	fetcher2.AssertExpectations(t)
}

func TestCollectAndExport_FetcherError(t *testing.T) {
	storage := new(mocks.Storage)
	fetcher1 := new(mocks.Fetcher)
	fetcher2 := new(mocks.Fetcher)

	data2 := []byte(`{"id":2}`)
	fetchError := errors.New("fetch failed")

	storage.On("Reset", mock.Anything).Return(nil)
	fetcher1.On("Fetch").Return(nil, fetchError)
	fetcher2.On("Fetch").Return(data2, nil)
	storage.On("Upload", mock.Anything, data2).Return(nil)

	uc := &usecase.UseCase{
		Config: usecase.Config{
			Timeout:    5 * time.Second,
			MaxWorkers: 2,
		},
		Storage: storage,
	}

	fetchers := []fetcher.Fetcher{fetcher1, fetcher2}
	errs := uc.CollectAndExport(fetchers)

	assert.NotNil(t, errs)
	assert.Len(t, errs, 1)
	assert.Equal(t, fetchError, errs[0])
	storage.AssertExpectations(t)
	fetcher1.AssertExpectations(t)
	fetcher2.AssertExpectations(t)
}

func TestCollectAndExport_UploadError(t *testing.T) {
	storage := new(mocks.Storage)
	fetcher1 := new(mocks.Fetcher)
	fetcher2 := new(mocks.Fetcher)

	data1 := []byte(`{"id":1}`)
	data2 := []byte(`{"id":2}`)
	uploadError := errors.New("upload failed")

	storage.On("Reset", mock.Anything).Return(nil)
	fetcher1.On("Fetch").Return(data1, nil)
	fetcher2.On("Fetch").Return(data2, nil)
	storage.On("Upload", mock.Anything, data1).Return(uploadError)
	storage.On("Upload", mock.Anything, data2).Return(nil)

	uc := &usecase.UseCase{
		Config: usecase.Config{
			Timeout:    5 * time.Second,
			MaxWorkers: 2,
		},
		Storage: storage,
	}

	fetchers := []fetcher.Fetcher{fetcher1, fetcher2}
	errs := uc.CollectAndExport(fetchers)

	assert.NotNil(t, errs)
	assert.Len(t, errs, 1)
	assert.Equal(t, uploadError, errs[0])
	storage.AssertExpectations(t)
	fetcher1.AssertExpectations(t)
	fetcher2.AssertExpectations(t)
}

func TestCollectAndExport_ResetError(t *testing.T) {
	storage := new(mocks.Storage)
	f := new(mocks.Fetcher)

	resetError := errors.New("reset failed")
	storage.On("Reset", mock.Anything).Return(resetError)

	uc := &usecase.UseCase{
		Config: usecase.Config{
			Timeout:    5 * time.Second,
			MaxWorkers: 1,
		},
		Storage: storage,
	}

	fetchers := []fetcher.Fetcher{f}
	errs := uc.CollectAndExport(fetchers)

	assert.NotNil(t, errs)
	assert.Len(t, errs, 1)
	storage.AssertExpectations(t)
	f.AssertNotCalled(t, "Fetch")
}

func TestCollectAndExport_Timeout(t *testing.T) {
	storage := new(mocks.Storage)
	f := new(mocks.Fetcher)

	storage.On("Reset", mock.Anything).Return(nil)
	f.On("Fetch").Run(func(args mock.Arguments) {
		time.Sleep(2 * time.Second)
	}).Return([]byte(`{"id":1}`), nil)

	uc := &usecase.UseCase{
		Config: usecase.Config{
			Timeout:    100 * time.Millisecond,
			MaxWorkers: 1,
		},
		Storage: storage,
	}

	fetchers := []fetcher.Fetcher{f}
	errs := uc.CollectAndExport(fetchers)

	assert.NotNil(t, errs)
	assert.Len(t, errs, 1)
	assert.Equal(t, context.DeadlineExceeded, errs[0])
	storage.AssertExpectations(t)
}

func TestCollectAndExport_MultipleFetchers(t *testing.T) {
	storage := new(mocks.Storage)

	numFetchers := 100
	fetchers := make([]fetcher.Fetcher, numFetchers)
	storage.On("Reset", mock.Anything).Return(nil)

	for i := 0; i < numFetchers; i++ {
		f := new(mocks.Fetcher)
		data := []byte(`{"id":` + string(rune(i)) + `}`)
		f.On("Fetch").Return(data, nil)
		storage.On("Upload", mock.Anything, data).Return(nil)
		fetchers[i] = f
	}

	uc := &usecase.UseCase{
		Config: usecase.Config{
			Timeout:    10 * time.Second,
			MaxWorkers: 10,
		},
		Storage: storage,
	}

	errs := uc.CollectAndExport(fetchers)
	assert.Nil(t, errs)
	storage.AssertExpectations(t)
}

func TestCollectAndExport_NoWorkerLimit(t *testing.T) {
	storage := new(mocks.Storage)
	fetcher1 := new(mocks.Fetcher)
	fetcher2 := new(mocks.Fetcher)

	data1 := []byte(`{"id":1}`)
	data2 := []byte(`{"id":2}`)

	storage.On("Reset", mock.Anything).Return(nil)
	fetcher1.On("Fetch").Return(data1, nil)
	fetcher2.On("Fetch").Return(data2, nil)
	storage.On("Upload", mock.Anything, data1).Return(nil)
	storage.On("Upload", mock.Anything, data2).Return(nil)

	uc := &usecase.UseCase{
		Config: usecase.Config{
			Timeout:    5 * time.Second,
			MaxWorkers: 0,
		},
		Storage: storage,
	}

	fetchers := []fetcher.Fetcher{fetcher1, fetcher2}
	errs := uc.CollectAndExport(fetchers)

	assert.Nil(t, errs)
	storage.AssertExpectations(t)
	fetcher1.AssertExpectations(t)
	fetcher2.AssertExpectations(t)
}

func TestCollectAndExport_EmptyFetchers(t *testing.T) {
	storage := new(mocks.Storage)
	storage.On("Reset", mock.Anything).Return(nil)

	uc := &usecase.UseCase{
		Config: usecase.Config{
			Timeout:    5 * time.Second,
			MaxWorkers: 2,
		},
		Storage: storage,
	}

	fetchers := []fetcher.Fetcher{}
	errs := uc.CollectAndExport(fetchers)

	assert.Nil(t, errs)
	storage.AssertExpectations(t)
}
