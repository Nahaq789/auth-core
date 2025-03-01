package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/auth-core/internal/domain/user"
	"github.com/auth-core/internal/infrastructure/mapper"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserRepositoryImpl struct {
	DynamoDBClient *dynamodb.Client
	tableName      string
}

const layout string = "2006-01-02 15:04:05"

func NewUserRepositoryImpl(client *dynamodb.Client, tableName string) *UserRepositoryImpl {
	return &UserRepositoryImpl{DynamoDBClient: client, tableName: tableName}
}

func (u *UserRepositoryImpl) Create(ctx context.Context, user *user.User) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	input := &dynamodb.PutItemInput{
		TableName: aws.String(u.tableName),
		Item: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberS{Value: user.UserId().Value()},
			"sub":     &types.AttributeValueMemberS{Value: user.Sub().Value()},
			"email":   &types.AttributeValueMemberS{Value: user.Email().Value()},
			"user_type": &types.AttributeValueMemberS{
				Value: user.UserType().String(),
			},
			"created_at": &types.AttributeValueMemberS{
				Value: user.CreatedAt().Format(layout),
			},
			"updated_at": &types.AttributeValueMemberS{
				Value: user.UpdatedAt().Format(layout),
			},
		},
	}

	_, err := u.DynamoDBClient.PutItem(ctx, input)
	if err != nil {
		if ctxErr := ctx.Err(); ctxErr == context.DeadlineExceeded {
			return fmt.Errorf("user create timeout: %s", ctxErr)
		}
		return err
	}
	return nil
}

func (u *UserRepositoryImpl) FindByUserId(ctx context.Context, userId user.UserId) (*user.User, error) {
	response, err := u.DynamoDBClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(u.tableName),
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberS{Value: userId.Value()},
		},
	})

	if err != nil {
		return nil, err
	}

	mapper := mapper.UserMapper{}
	user, err := mapper.MapToDomain(response)
	if err != nil {
		return nil, err
	}
	return user, nil
}
