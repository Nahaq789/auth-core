package repository

import (
	"context"

	"github.com/auth-core/internal/domain/models/auth"
	valueObjects "github.com/auth-core/internal/domain/value_objects"
)

type CognitoRepository interface {
	SignUp(ctx context.Context, user *auth.SignUp) (*auth.SignUpResult, error)
	ConfirmSignUp(ctx context.Context, c *auth.ConfirmSignUp) error
	InitiateAuth(ctx context.Context, s *auth.Credentials) (*valueObjects.AuthenticationChallenge, error)
	AuthChallenge(ctx context.Context, a *auth.AuthChallenge) (*valueObjects.Token, error)
}
