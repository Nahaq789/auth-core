package dto

type SignInDto struct {
	Email string `json:"email" binding:"required"`
	SrpA  string `json:"srp_a" binding:"required"`
}
