package user

import (
	"fmt"
	"strings"

	"github.com/auth-core/internal/domain"
)

type UserId struct {
	value string
}

const Prefix = "usr"

func (u UserId) Value() string {
	return u.value
}

func (u UserId) String() string {
	return u.value
}

func NewUserId(uuid domain.Uuid) (*UserId, error) {
	v, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UserId: %w", err)
	}
	userId := fmt.Sprintf("%s_%s", Prefix, v)

	return &UserId{value: userId}, nil
}

func UserIdFromStr(userId string) (*UserId, error) {
	if err := validatePrefix(userId); err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &UserId{value: userId}, nil
}

func validatePrefix(userId string) error {
	if !strings.HasPrefix(userId, Prefix) {
		return fmt.Errorf("invalid user id format: expected prefix '%s', got '%s'", Prefix, userId)
	}

	return nil
}
