package services

import (
	"context"
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
		c.logger.Error("Faild email", "email", d.Email, "error", err)
		return err
	}
	password := valueObjects.NewPassword(d.Password)

	auth := auth.NewAuth(
		*email,
		*password,
	)

	result, err := c.cognito.SignUp(ctx, auth)
	if err != nil {
		c.logger.Error("Failed Signup user",
			"email", auth.Email().String(),
			"error", err,
		)
		return err
	}

	var uuidImpl uuid.UuidImpl
	u, err := uuidImpl.NewV4()
	if err != nil {
		c.logger.Error("Failed generate uuid", "error", err)
		return err
	}
	userId, err := user.NewUserId(uuid.NewUuid(u))
	if err != nil {
		c.logger.Error("Failed generate userId", "error", err)
		return err
	}

	userType := user.NewUserType("standard")
	time := time.Now()
	user := user.NewUser(*userId, *result.Sub, *email, userType, time, time)

	c.logger.Info("Complete SignUp user", "email", auth.Email().String())
	c.logger.Info("Finish Cognito SignUp")
	return nil
}
