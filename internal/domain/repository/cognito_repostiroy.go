package repository

import (
	"context"

	"github.com/auth-core/internal/domain/auth"
)

type CognitoRepository interface {
	SignUp(ctx context.Context, user *auth.Auth) (bool, error)
}
