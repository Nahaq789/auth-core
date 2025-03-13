package services

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/auth-core/internal/application"
	"github.com/auth-core/internal/application/dto"
	"github.com/auth-core/internal/domain/models/auth"
	"github.com/auth-core/internal/domain/models/user"
	"github.com/auth-core/internal/domain/repository"
	valueObjects "github.com/auth-core/internal/domain/value_objects"
	"github.com/auth-core/pkg/uuid"
)

type CognitoServiceImpl struct {
	logger      *slog.Logger
	userService application.UserService
	cognito     repository.CognitoRepository
}

func NewCognitoService(
	logger *slog.Logger,
	userService application.UserService,
	cognito repository.CognitoRepository,
) *CognitoServiceImpl {
	return &CognitoServiceImpl{logger: logger, userService: userService, cognito: cognito}
}

func (c *CognitoServiceImpl) SignUp(ctx context.Context, d *dto.AuthDto) error {
	c.logger.Info("Start Cognito SignUp", "email", d.Email)

	email, err := valueObjects.NewEmail(d.Email)
	if err != nil {
		c.logger.Error("Failed email validation", "email", d.Email, "error", err)
		return fmt.Errorf("invalid email format %q: %w", d.Email, err)
	}
	password := valueObjects.NewPassword(d.Password)

	auth := auth.NewAuth(
		*email,
		*password,
	)

	result, err := c.cognito.SignUp(ctx, auth)
	if err != nil {
		c.logger.Error("Failed Cognito SignUp",
			"email", auth.Email().String(),
			"error", err,
		)
		return fmt.Errorf("cognito signup failed for email %q: %w", auth.Email().String(), err)
	}

	var uuidImpl uuid.UuidImpl
	userId, err := user.NewUserId(uuidImpl)
	if err != nil {
		c.logger.Error("Failed to generate userId", "error", err)
		return fmt.Errorf("failed to generate user ID: %w", err)
	}

	userType := user.NewUserType("standard")
	time := time.Now()
	user := user.NewUser(*userId, *result.Sub, *email, userType, time, time)
	err = c.userService.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	c.logger.Info("Complete SignUp user",
		"email", auth.Email().String(),
		"userId", userId.Value())
	c.logger.Info("Finish Cognito SignUp")
	return nil
}
