package port

type Fetcher[T any] interface {
	Fetch(cfg T) ([]byte, error)
}
