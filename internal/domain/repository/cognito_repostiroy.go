package repository

import (
	"context"

	"github.com/auth-core/internal/domain/models/auth"
)

type CognitoRepository interface {
	SignUp(ctx context.Context, user *auth.Auth) (*auth.SignUpResult, error)
	VerifyCode(ctx context.Context, c *auth.VerifyCode) error
}
