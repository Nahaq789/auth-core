package application

import (
	"context"

	"github.com/auth-core/internal/application/dto"
	"github.com/auth-core/internal/domain/user"
)

type UserService interface {
	CreateUser(ctx context.Context, d *dto.UserDto) error
	FindByUserId(ctx context.Context, userId string) (*user.User, error)
}

type CognitoService interface {
	SignUp(ctx context.Context, d *dto.AuthDto) error
}
