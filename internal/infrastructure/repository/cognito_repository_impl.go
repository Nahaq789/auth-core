package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/auth-core/internal/domain/models/auth"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type CognitoRepositoryImpl struct {
	CognitoClient *cognitoidentityprovider.Client
	clientId      string
}

func NewCognitoRepository(client *cognitoidentityprovider.Client, clientId string) *CognitoRepositoryImpl {
	return &CognitoRepositoryImpl{
		CognitoClient: client,
		clientId:      clientId,
	}
}

func (actor *CognitoRepositoryImpl) SignUp(ctx context.Context, a *auth.Auth) (*auth.SignUpResult, error) {
	output, err := actor.CognitoClient.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(actor.clientId),
		Password: aws.String(a.Password().String()),
		Username: aws.String(a.Email().String()),
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(a.Email().String())},
		},
	})

	if err != nil {
		var invalidPassword *types.InvalidPasswordException
		if errors.As(err, &invalidPassword) {
			return nil, fmt.Errorf("%s", *invalidPassword.Message)
		} else {
			return nil, fmt.Errorf("Couldn't sign up user %v. message: %w\n", a.Email().String(), err)
		}
	}

	result, err := auth.NewSignUpResult(*output.UserSub)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return *&result, err
}

func (actor *CognitoRepositoryImpl) VerifyCode(ctx context.Context, c *auth.VerifyCode) error {
	output, err := actor.CognitoClient.VerifyUserAttribute(ctx, &cognitoidentityprovider.VerifyUserAttributeInput{
		AttributeName: aws.String("email"),
		Code:          aws.String(c.Code()),
	})
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	fmt.Println(output)
	return nil
}
