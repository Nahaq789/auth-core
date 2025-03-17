package auth

import valueObjects "github.com/auth-core/internal/domain/value_objects"

type SignUp struct {
	email    valueObjects.Email
	password valueObjects.Password
}

func NewSignUp(email valueObjects.Email, password valueObjects.Password) *SignUp {
	return &SignUp{email: email, password: password}
}

func (a *SignUp) Email() valueObjects.Email {
	return a.email
}

func (a *SignUp) Password() valueObjects.Password {
	return a.password
}
