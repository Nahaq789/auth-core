package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func NewCognitoClient(ctx context.Context) (*cognitoidentityprovider.Client, error) {
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return nil, err
	}

	cognitoClient := cognitoidentityprovider.NewFromConfig(sdkConfig)

	return cognitoClient, nil
}
