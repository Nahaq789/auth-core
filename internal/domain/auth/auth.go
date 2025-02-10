package auth

import valueObjects "github.com/auth-core/internal/domain/value_objects"

type Auth struct {
	email valueObjects.Email
}

func NewAuth(email valueObjects.Email) *Auth {
	return &Auth{email: email}
}
