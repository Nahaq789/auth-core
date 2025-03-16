package auth

import (
	valueObjects "github.com/auth-core/internal/domain/value_objects"
)

type ConfirmSignUp struct {
	userName valueObjects.Email
	code     string
}

func NewConfirmSignUp(userName valueObjects.Email, code string) *ConfirmSignUp {
	return &ConfirmSignUp{userName: userName, code: code}
}

func (c *ConfirmSignUp) Code() string {
	return c.code
}

func (c *ConfirmSignUp) UserName() valueObjects.Email {
	return c.userName
}
