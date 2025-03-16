package dto

type ConfirmSignUpDto struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

func NewConfirmSignUpDto(email string, code string) *ConfirmSignUpDto {
	return &ConfirmSignUpDto{Email: email, Code: code}
}
