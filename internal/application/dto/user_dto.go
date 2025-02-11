package dto

import "time"

type UserDto struct {
	userId    string
	sub       string
	email     string
	userType  string
	createdAt time.Time
	updatedAt time.Time
}

func (u UserDto) UserId() string {
	return u.userId
}

func (u UserDto) Sub() string {
	return u.sub
}

func (u UserDto) Email() string {
	return u.email
}

func (u UserDto) UserType() string {
	return u.userType
}

func (u UserDto) CreatedAt() time.Time {
	return u.createdAt
}

func (u UserDto) UpdatedAt() time.Time {
	return u.updatedAt
}

func NewUserDto(userId string, sub string, email string, userType string, createdAt time.Time, updatedAt time.Time) *UserDto {
	return &UserDto{
		userId:    userId,
		sub:       sub,
		email:     email,
		userType:  userType,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
