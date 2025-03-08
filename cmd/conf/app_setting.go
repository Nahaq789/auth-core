package conf

import (
	"context"
	"fmt"

	"github.com/auth-core/pkg/db"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type AppSetting struct {
	Server ServerSetting
	Aws    AwsSetting
}

type AwsSetting struct {
	Region          string `env:"REGION"`
	UserTable       string `env:"USER_TABLE"`
	CognitoClientId string `env:"CLIENT_ID"`
}

type AwsClient struct {
	Dynamodb dynamodb.Client
	Cognito  cognitoidentityprovider.Client
}

func (a *AwsClient) InitClient(ctx context.Context, aws *AwsSetting) (*AwsClient, error) {
	dynamodb, err := db.NewDynamoDbClient(ctx, aws.Region)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}
	cognito := cognitoidentityprovider.NewFromConfig(sdkConfig)

	awsClient := AwsClient{
		Dynamodb: *dynamodb,
		Cognito:  *cognito,
	}
	return &awsClient, nil
}

type ServerSetting struct {
	Port  string `env:"PORT"`
	Level string `env:"LEVEL"`
}
