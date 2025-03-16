package dto

type VerifyCodeDto struct {
	Code string `json:"code" binding:"required"`
}

func NewVerifyCodeDto(code string) *VerifyCodeDto {
	return &VerifyCodeDto{Code: code}
}
