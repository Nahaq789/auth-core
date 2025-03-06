package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/auth-core/internal/domain/auth"
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

func (actor *CognitoRepositoryImpl) SignUp(ctx context.Context, auth *auth.Auth) (bool, error) {
	output, err := actor.CognitoClient.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(actor.clientId),
		Password: aws.String(auth.Password().String()),
		Username: aws.String(auth.Email().String()),
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(auth.Email().String())},
		},
	})

	if err != nil {
		var invalidPassword *types.InvalidPasswordException
		if errors.As(err, &invalidPassword) {
			return output.UserConfirmed, fmt.Errorf("%s", *invalidPassword.Message)
		} else {
			return output.UserConfirmed, fmt.Errorf("Couldn't sign up user %v. message: %v\n", auth.Email().String(), err)
		}
	}

	return output.UserConfirmed, err
}
