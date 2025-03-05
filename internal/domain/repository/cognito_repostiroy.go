package repository

import (
	"context"
	"github.com/auth-core/internal/domain/user"
)

type CognitoRepository interface {
	SignUp(ctx context.Context, user *user.User) (bool, error)
}
