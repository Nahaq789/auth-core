package auth

import "github.com/auth-core/internal/domain"

type SignUpResult struct {
	Sub string
}

func NewSignUpResult(sub string) (*SignUpResult, error) {
	err := domain.Uuid.ValidFormat(sub)
	if err != nil {
		return nil, err
	}
	return &SignUpResult{Sub: sub}, err
}
