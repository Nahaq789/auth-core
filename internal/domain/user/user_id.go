package user

import (
	"fmt"

	"github.com/auth-core/internal/domain"
)

type UserId struct {
	value string
}

func (u UserId) Value() string {
	return u.value
}

func (u UserId) String() string {
	return fmt.Sprintf("%v", u.value)
}

func NewUserId(uuid domain.Uuid) (*UserId, error) {
	v, err := uuid.NewV4()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UserId: %w", err)
	}
	return &UserId{value: v}, nil
}
