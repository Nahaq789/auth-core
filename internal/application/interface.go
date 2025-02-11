package application

import (
	"context"

	"github.com/auth-core/internal/domain/user"
)

type UserService interface {
	Create(ctx context.Context, user *user.User) error
	FindByUserId(ctx context.Context, userId user.UserId) (*user.User, error)
}
