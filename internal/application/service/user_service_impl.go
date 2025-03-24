package services

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/auth-core/internal/domain/models/user"
	"github.com/auth-core/internal/domain/repository"
)

type UserServiceImpl struct {
	logger     *slog.Logger
	repository repository.UserRepository
	cognito    repository.AuthRepository
}

func NewUserService(logger *slog.Logger, repository repository.UserRepository, cognito repository.AuthRepository) *UserServiceImpl {
	return &UserServiceImpl{logger: logger, repository: repository, cognito: cognito}
}

func (us *UserServiceImpl) CreateUser(ctx context.Context, u *user.User) error {
	us.logger.Info("Start Create User", "email", u.Email(), "sub", u.Sub())

	exist, err := us.repository.Exist(ctx, u.Email().Value())
	if err != nil {
		us.logger.Error("Failed to check if user exists", "email", u.Email().Value(), "error", err)
		return fmt.Errorf("failed to check if user with email %q exists: %w", u.Email().Value(), err)
	}
	if exist {
		us.logger.Error("user already exist", "error", u.Email().Value())
		return fmt.Errorf("user with email %q already exists", u.Email().Value())
	}

	err = us.repository.Create(ctx, u)
	if err != nil {
		us.logger.Error("Failed to create user", "email", u.Email().String(), "error", err)
		return fmt.Errorf("failed to save user with email %q to database: %w", u.Email().String(), err)
	}

	us.logger.Info("user created", "email", u.Email().Value())
	us.logger.Info("Finish Create User")
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
