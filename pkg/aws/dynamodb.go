package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoDbClient(ctx context.Context, region string) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("can not load config: %s", err)
	}

	client := dynamodb.NewFromConfig(cfg)

	return client, nil
}
