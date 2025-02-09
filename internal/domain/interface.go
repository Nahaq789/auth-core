package domain

type Uuid interface {
	NewV4() (string, error)
}
