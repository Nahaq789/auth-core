package repository

import (
	"context"

	"github.com/auth-core/internal/domain/models/user"
)

type UserRepository interface {
	Create(ctx context.Context, user *user.User) error
	FindByUserId(ctx context.Context, userId user.UserId) (*user.User, error)
	Exist(ctx context.Context, email string) (bool, error)
}
