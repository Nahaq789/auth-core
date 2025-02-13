package valueObjects

type Pasawoed struct {
	value string
}

func NewPassword(value string) *Pasawoed {
	return &Pasawoed{value: value}
}
