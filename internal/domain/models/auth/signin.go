package auth

import valueObjects "github.com/auth-core/internal/domain/value_objects"

type SignIn struct {
	email    valueObjects.Email
	password valueObjects.Password
}

func NewSignIn(email valueObjects.Email, password valueObjects.Password) *SignIn {
	return &SignIn{email: email, password: password}
}

func (s *SignIn) Email() valueObjects.Email {
	return s.email
}

func (s *SignIn) Password() valueObjects.Password {
	return s.password
}
