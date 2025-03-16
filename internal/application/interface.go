package application

import (
	"context"

	"github.com/auth-core/internal/application/dto"
	"github.com/auth-core/internal/domain/models/user"
)

type UserService interface {
	CreateUser(ctx context.Context, user *user.User) error
	FindByUserId(ctx context.Context, userId string) (*user.User, error)
}

type CognitoService interface {
	SignUp(ctx context.Context, d *dto.AuthDto) error
	VerifyCode(ctx context.Context, code *dto.VerifyCodeDto) error
}
