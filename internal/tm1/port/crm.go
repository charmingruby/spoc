package port

type CRM interface {
	Authenticate(apiKey string, shouldSimulateError bool) (string, error)
	GenerateRelatory(token string, shouldSimulateError bool) ([]byte, error)
}
