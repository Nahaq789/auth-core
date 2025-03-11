package domain

type Uuid interface {
	NewV4() (string, error)
	ValidFormat(uuid string) error
}
