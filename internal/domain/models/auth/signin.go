package auth

import valueObjects "github.com/auth-core/internal/domain/value_objects"

type SignIn struct {
	email valueObjects.Email
	srpA  string
}

func NewSignIn(email valueObjects.Email, password string) *SignIn {
	return &SignIn{email: email, srpA: password}
}

func (s *SignIn) Email() valueObjects.Email {
	return s.email
}

func (s *SignIn) SrpA() string {
	return s.srpA
}
