package valueObjects

import (
	"fmt"
	"net/mail"
)

type Email struct {
	value string
}

func (e Email) Value() string {
	return e.value
}

func (e Email) String() string {
	return e.value
}

func NewEmail(email string) (*Email, error) {
	if err := validateEmail(email); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &Email{value: email}, nil
}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("invalid email address: %w", err)
	}

	return nil
}
