package fetcher

type Fetcher interface {
	Fetch() ([]byte, error)
}
