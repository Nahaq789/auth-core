package dto

type SignInDto struct {
	email string `json:"email" binding:"required"`
	srpA  string `json:"srp_a" binding:"required"`
}

func (s *SignInDto) Email() string {
	return s.email
}

func (s *SignInDto) SrpA() string {
	return s.srpA
}
