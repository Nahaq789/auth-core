package services

import (
	"context"
	"errors"
	"log/slog"

	"github.com/auth-core/internal/application/dto"
	"github.com/auth-core/internal/domain/repository"
	"github.com/auth-core/internal/domain/user"
	valueObjects "github.com/auth-core/internal/domain/value_objects"
)

type UserServiceImpl struct {
	logger     *slog.Logger
	repository repository.UserRepository
	cognito    repository.CognitoRepository
}

func NewUserService(logger *slog.Logger, repository repository.UserRepository, cognito repository.CognitoRepository) *UserServiceImpl {
	return &UserServiceImpl{logger: logger, repository: repository, cognito: cognito}
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, d *dto.UserDto) error {
	exist, err := u.repository.Exist(ctx, d.Email)
	if err != nil {
		u.logger.Error("Failed to check if user exists", "email", d.Email, "error", err)
		return err
	}
	if exist {
		u.logger.Error("user already exist", "error", d.Email)
		return errors.New("user with this email already exist")
	}

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

	u.logger.Info("user created", "email", d.Email)
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
