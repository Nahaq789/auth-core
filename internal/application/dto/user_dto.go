package dto

import "time"

type UserDto struct {
	userId    string    `json:"user_id" binding:"required"`
	sub       string    `json:"sub" binding:"required"`
	email     string    `json:"email" binding:"required"`
	userType  string    `json:"user_type" binding:"required"`
	createdAt time.Time `json:"created_at" binding:"required"`
	updatedAt time.Time `json:"updated_at" binding:"required"`
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
