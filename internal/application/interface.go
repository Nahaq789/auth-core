package application

import (
	"context"

	"github.com/auth-core/internal/application/dto"
	"github.com/auth-core/internal/domain/user"
)

type UserService interface {
	Create(ctx context.Context, user *dto.UserDto) error
	FindByUserId(ctx context.Context, userId string) (*user.User, error)
}
