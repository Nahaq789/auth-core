package mapper

import (
	"fmt"
	"time"

	"github.com/auth-core/internal/domain/user"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserMapper struct{}

const layout = "2006-01-02 15:04:05"

func (u *UserMapper) MapToDomain(item *dynamodb.GetItemOutput) (*user.User, error) {
	idAttr, err := getStringValue(item.Item, "user_id")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	subAttr, err := getStringValue(item.Item, "sub")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	emailAttr, err := getStringValue(item.Item, "email")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	userTypeAttr, err := getStringValue(item.Item, "user_type")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	createdAtAttr, err := getStringValue(item.Item, "created_at")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	updatedAtAttr, err := getStringValue(item.Item, "updated_at")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return generateUser(idAttr, subAttr, emailAttr, userTypeAttr, createdAtAttr, updatedAtAttr)
}

func generateUser(idAttr string, subAttr string, emailAttr string, userTypeAttr string, createdAtAttr string, updatedAtAttr string) (*user.User, error) {
	userId, err := user.UserIdFromStr(idAttr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user_id: %w", err)
	}
	sub := user.NewSub(subAttr)
	email, err := user.NewEmail(emailAttr)
	if err != nil {
		return nil, fmt.Errorf("failed to parsee email: %w", err)
	}
	userType := user.NewUserType(userTypeAttr)
	createdAt, err := timeParse(createdAtAttr)
	if err != nil {
		return nil, fmt.Errorf("failed to parsee created_at: %w", err)
	}
	updatedAt, err := timeParse(updatedAtAttr)
	if err != nil {
		return nil, fmt.Errorf("failed to parsee updated_at: %w", err)
	}

	return user.NewUser(*userId, *sub, *email, userType, createdAt, updatedAt), nil
}

func getStringValue(item map[string]types.AttributeValue, key string) (string, error) {
	attr, ok := item[key]
	if !ok {
		return "", fmt.Errorf("key %s not found", key)
	}

	attrStr, ok := attr.(*types.AttributeValueMemberS)
	if !ok {
		return "", fmt.Errorf("key %s is not string", key)
	}

	return attrStr.Value, nil
}

func timeParse(value string) (time.Time, error) {
	return time.Parse(layout, value)
}
