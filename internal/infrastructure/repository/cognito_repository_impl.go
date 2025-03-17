package repository

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/auth-core/internal/domain/models/auth"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

var AuthFlowTypeUserSrpAuth = "USER_SRP_AUTH"

type CognitoRepositoryImpl struct {
	CognitoClient *cognitoidentityprovider.Client
	clientId      string
	clientSecret  string
}

func NewCognitoRepository(client *cognitoidentityprovider.Client, clientId string, clientSecret string) *CognitoRepositoryImpl {
	return &CognitoRepositoryImpl{
		CognitoClient: client,
		clientId:      clientId,
		clientSecret:  clientSecret,
	}
}

func (actor *CognitoRepositoryImpl) SignUp(ctx context.Context, a *auth.SignUp) (*auth.SignUpResult, error) {
	secretHash := actor.generateSecretHash(a.Email().Value())
	output, err := actor.CognitoClient.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
		ClientId:   aws.String(actor.clientId),
		Password:   aws.String(a.Password().String()),
		Username:   aws.String(a.Email().String()),
		SecretHash: &secretHash,
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(a.Email().String())},
		},
	})

	if err != nil {
		var invalidPassword *types.InvalidPasswordException
		if errors.As(err, &invalidPassword) {
			return nil, fmt.Errorf("%s", *invalidPassword.Message)
		} else {
			return nil, fmt.Errorf("Failed to sign up user %v. message: %w\n", a.Email().String(), err)
		}
	}

	result, err := auth.NewSignUpResult(*output.UserSub)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return *&result, err
}

func (actor *CognitoRepositoryImpl) ConfirmSignUp(ctx context.Context, c *auth.ConfirmSignUp) error {
	secretHash := actor.generateSecretHash(c.UserName().Value())
	code := c.Code()
	userName := c.UserName().Value()
	output, err := actor.CognitoClient.ConfirmSignUp(ctx, &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         &actor.clientId,
		ConfirmationCode: &code,
		Username:         &userName,
		SecretHash:       &secretHash,
	})
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	fmt.Println(output)
	return nil
}

func (actor *CognitoRepositoryImpl) SignIn(ctx context.Context, s *auth.SignIn) error {
	output, err := actor.CognitoClient.InitiateAuth(ctx, &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       types.AuthFlowType(AuthFlowTypeUserSrpAuth),
		ClientId:       aws.String(actor.clientId),
		AuthParameters: map[string]string{"USERNAME": s.Email().Value(), "SRP_A": s.SrpA()},
	})
	if err != nil {
		var invalidPassword *types.InvalidPasswordException
		if errors.As(err, &invalidPassword) {
			return fmt.Errorf("%s", *invalidPassword.Message)
		} else {
			return fmt.Errorf("Failed to signin user %v. message: %w\n", s.Email().String(), err)
		}
	}
	fmt.Println(output)

	return nil
}

func (actor *CognitoRepositoryImpl) generateSecretHash(userName string) string {
	mac := hmac.New(sha256.New, []byte(actor.clientSecret))
	mac.Write([]byte(userName + actor.clientId))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
