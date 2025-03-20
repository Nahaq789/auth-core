package auth

import valueObjects "github.com/auth-core/internal/domain/value_objects"

type Credentials struct {
	email valueObjects.Email
	srpA  string
}

func NewCredentials(email valueObjects.Email, password string) *Credentials {
	return &Credentials{email: email, srpA: password}
}

func (s *Credentials) Email() valueObjects.Email {
	return s.email
}

func (s *Credentials) SrpA() string {
	return s.srpA
}
