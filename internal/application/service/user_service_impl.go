package services

import (
	"context"

	"github.com/auth-core/internal/application/dto"
	"github.com/auth-core/internal/domain/repository"
	"github.com/auth-core/internal/domain/user"
	valueObjects "github.com/auth-core/internal/domain/value_objects"
)

type UserServiceImpl struct {
	repository repository.UserRepository
	cognito    repository.CognitoRepository
}

func NewUserService(repository repository.UserRepository, cognito repository.CognitoRepository) *UserServiceImpl {
	return &UserServiceImpl{repository: repository, cognito: cognito}
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, d *dto.UserDto) error {
	userId, err := user.UserIdFromStr(d.UserId)
	if err != nil {
		return err
	}
	sub := user.NewSub(d.Sub)
	email, err := valueObjects.NewEmail(d.Email)
	if err != nil {
		return err
	}
	userType := user.NewUserType(d.UserType)

	user := user.NewUser(
		*userId, *sub, *email, userType, d.CreatedAt, d.UpdatedAt)

	err = u.repository.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserServiceImpl) FindByUserId(ctx context.Context, userId string) (*user.User, error) {
	id, err := user.UserIdFromStr(userId)
	if err != nil {
		return nil, err
	}

	user, err := u.repository.FindByUserId(ctx, *id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
