package services

import (
	"context"

	"github.com/auth-core/internal/domain/repository"
	"github.com/auth-core/internal/domain/user"
)

type UserServiceImpl struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserServiceImpl {
	return UserServiceImpl{repository: repository}
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, user user.User) error {
	return nil
}

func (u *UserServiceImpl) FindByUserId(ctx context.Context, userId user.UserId) error {
	return nil
}
