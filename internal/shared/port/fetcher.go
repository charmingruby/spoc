package port

type Fetcher interface {
	Fetch() ([]byte, error)
}
