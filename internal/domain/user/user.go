package user

import (
	"time"

	valueObjects "github.com/auth-core/internal/domain/value_objects"
)

type User struct {
	userId    UserId
	sub       Sub
	email     valueObjects.Email
	userType  UserType
	createdAt time.Time
	updatedAt time.Time
}

func (u User) UserId() UserId {
	return u.userId
}

func (u User) Sub() Sub {
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

func NewUser(userId UserId, sub Sub, email valueObjects.Email, userType UserType, createdAt time.Time, updatedAt time.Time) *User {
	return &User{
		userId:    userId,
		sub:       sub,
		email:     email,
		userType:  userType,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
