package port

type TM3 interface {
	GenerateRelatory(shouldSimulateError bool) ([]byte, error)
}
