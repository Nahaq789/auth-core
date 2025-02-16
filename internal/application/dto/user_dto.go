package dto

import "time"

type UserDto struct {
	UserId    string `json:"user_id" binding:"required"`
	Sub       string `json:"sub" binding:"required"`
	Email     string `json:"email" binding:"required"`
	UserType  string `json:"user_type" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserDto(userId string, sub string, email string, userType string, createdAt time.Time, updatedAt time.Time) *UserDto {
	return &UserDto{
		UserId:    userId,
		Sub:       sub,
		Email:     email,
		UserType:  userType,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
