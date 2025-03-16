package repository

import (
	"context"

	"github.com/auth-core/internal/domain/models/auth"
)

type CognitoRepository interface {
	SignUp(ctx context.Context, user *auth.Auth) (*auth.SignUpResult, error)
	ConfirmSignUp(ctx context.Context, c *auth.ConfirmSignUp) error
	SignIn(ctx context.Context, s *auth.SignIn) error
}
