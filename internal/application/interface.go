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
	SignUp(ctx context.Context, d *dto.SignUpDto) error
	ConfirmSignUp(ctx context.Context, code *dto.ConfirmSignUpDto) error
	InitiateAuth(ctx context.Context, d *dto.SignInDto) (*dto.InitiateAuthResultDto, error)
}
