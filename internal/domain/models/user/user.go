package user

import (
	"fmt"
	"time"

	"github.com/auth-core/internal/domain/models/sub"
	valueObjects "github.com/auth-core/internal/domain/value_objects"
)

const layout = "2006-01-02 15:04:05"

type User struct {
	userId    UserId
	sub       sub.Sub
	email     valueObjects.Email
	userType  UserType
	createdAt time.Time
	updatedAt time.Time
}

func (u User) UserId() UserId {
	return u.userId
}

func (u User) Sub() sub.Sub {
	return u.sub
}

func (u User) Email() valueObjects.Email {
	return u.email
}

func (u User) UserType() UserType {
	return u.userType
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

func (u User) UpdatedAt() time.Time {
	return u.updatedAt
}

func NewUser(userId UserId, sub sub.Sub, email valueObjects.Email, userType UserType, createdAt time.Time, updatedAt time.Time) *User {
	return &User{
		userId:    userId,
		sub:       sub,
		email:     email,
		userType:  userType,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func ParseTime(t string) (*time.Time, error) {
	time, err := time.Parse(layout, t)
	if err != nil {
		return nil, fmt.Errorf("can not parse time %s", t)
	}
	return &time, nil
}
