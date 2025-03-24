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
	cognito     repository.AuthRepository
}

func NewCognitoService(
	logger *slog.Logger,
	userService application.UserService,
	cognito repository.AuthRepository,
) *CognitoServiceImpl {
	return &CognitoServiceImpl{logger: logger, userService: userService, cognito: cognito}
}

func (c *CognitoServiceImpl) SignUp(ctx context.Context, d *dto.SignUpDto) error {
	c.logger.Info("Start Cognito SignUp", "email", d.Email)

	email, err := valueObjects.NewEmail(d.Email)
	if err != nil {
		c.logger.Error("Failed email validation", "email", d.Email, "error", err)
		return fmt.Errorf("invalid email format %q: %w", d.Email, err)
	}
	password := valueObjects.NewPassword(d.Password)

	auth := auth.NewSignUp(
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

func (c *CognitoServiceImpl) ConfirmSignUp(ctx context.Context, confirm *dto.ConfirmSignUpDto) error {
	email, err := valueObjects.NewEmail(confirm.Email)
	if err != nil {
		c.logger.Error("Failed to email",
			"email", email,
			"error", err)
		return fmt.Errorf("invalid email format %q: %w", confirm.Email, err)

	}
	verifyCode := auth.NewConfirmSignUp(*email, confirm.Code)
	if err := c.cognito.ConfirmSignUp(ctx, verifyCode); err != nil {
		c.logger.Error("Failed to ConfirmSignUp", "error", err)
		return fmt.Errorf("failed to confirm signup: %w", err)
	}

	return nil
}

func (c *CognitoServiceImpl) InitiateAuth(ctx context.Context, d *dto.SignInDto) (*dto.InitiateAuthResultDto, error) {
	email, err := valueObjects.NewEmail(d.Email)
	if err != nil {
		c.logger.Error("Failed email validation", "email", d.Email, "error", err)
		return nil, fmt.Errorf("invalid email format %q: %w", d.Email, err)
	}

	signin := auth.NewCredentials(*email, d.SrpA)
	challenge, err := c.cognito.InitiateAuth(ctx, signin)
	if err != nil {
		c.logger.Error("Failed to InitiateAuth", "error", err)
		return nil, fmt.Errorf("failed to initiate auth: %w", err)
	}

	return &dto.InitiateAuthResultDto{
		ChallengeName: challenge.GetChallengeName(),
		SrpB:          challenge.GetSrpB(),
		Salt:          challenge.GetSalt(),
		SecretBlock:   challenge.GetSecretBlock(),
		UserIdForSrp:  challenge.GetUserIdForSrp(),
	}, nil
}

func (c *CognitoServiceImpl) AuthChallenge(ctx context.Context, d *dto.AuthChallengeDto) (*dto.AuthChallengeResultDto, error) {
	email, err := valueObjects.NewEmail(d.Email)
	if err != nil {
		c.logger.Error("Failed email validation", "email", d.Email, "error", err)
		return nil, fmt.Errorf("invalid email format %q: %w", d.Email, err)
	}
	challenge := auth.NewAuthChallenge(d.TimeStamp, *email, d.SecretBlock, d.Signature)

	token, err := c.cognito.AuthChallenge(ctx, challenge)
	if err != nil {
		c.logger.Error("Failed to AuthChallenge", "error", err)
		return nil, fmt.Errorf("failed to auth challenge: %w", err)
	}
	fmt.Println(token)
	return &dto.AuthChallengeResultDto{
		AccessToken:  token.AccessToken(),
		IdToken:      token.IdToken(),
		RefreshToken: token.RefreshToken(),
	}, nil
}
