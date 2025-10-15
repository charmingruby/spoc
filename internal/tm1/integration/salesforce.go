package integration

type Salesforce interface {
	Authenticate(apiKey string, shouldSimulateError bool) (string, error)
	GenerateRelatory(token string, shouldSimulateError bool) ([]byte, error)
}
