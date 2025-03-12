package auth

import (
	"github.com/auth-core/internal/domain/models/sub"
)

type SignUpResult struct {
	Sub *sub.Sub
}

func NewSignUpResult(s string) (*SignUpResult, error) {
	sub, err := sub.NewSub(s)
	if err != nil {
		return nil, err
	}
	return &SignUpResult{Sub: sub}, nil
}
