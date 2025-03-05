package valueObjects

type Password struct {
	value string
}

func (p Password) Value() string {
	return p.value
}

func (p Password) String() string {
	return p.value
}

func NewPassword(value string) *Password {
	return &Password{value: value}
}
