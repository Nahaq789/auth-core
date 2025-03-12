package auth

import valueObjects "github.com/auth-core/internal/domain/value_objects"

type Auth struct {
	email    valueObjects.Email
	password valueObjects.Password
}

func NewAuth(email valueObjects.Email, password valueObjects.Password) *Auth {
	return &Auth{email: email, password: password}
}

func (a *Auth) Email() valueObjects.Email {
	return a.email
}

func (a *Auth) Password() valueObjects.Password {
	return a.password
}
