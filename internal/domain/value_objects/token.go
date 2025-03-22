package valueObjects

import "fmt"

type Token struct {
	accessToken  string
	idToken      string
	refreshToken string
}

func NewToken(a string, i string, r string) (*Token, error) {
	err := validateToken(a, i, r)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &Token{
		accessToken:  a,
		idToken:      i,
		refreshToken: r,
	}, nil
}

func validateToken(a string, i string, r string) error {
	if a == "" || i == "" || r == "" {
		return fmt.Errorf("need all tokens.")
	}
	return nil
}

func (t *Token) AccessToken() string {
	return t.accessToken
}

func (t *Token) IdToken() string {
	return t.idToken
}

func (t *Token) RefreshToken() string {
	return t.refreshToken
}
