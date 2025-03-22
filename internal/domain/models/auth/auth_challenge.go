package auth

import valueObjects "github.com/auth-core/internal/domain/value_objects"

const CHALLENGE_NAME = "PASSWORD_VERIFIER"

type AuthChallenge struct {
	challengeName     string
	timeStamp         string
	email             valueObjects.Email
	secretBlock       string
	signature         string
	challengeResponse map[string]string
}

func NewAuthChallenge(
	timeStamp string,
	userName valueObjects.Email,
	secretBlock string,
	signature string,
) *AuthChallenge {
	res := map[string]string{
		"TIMESTAMP":                   timeStamp,
		"USERNAME":                    userName.String(),
		"PASSWORD_CLAIM_SECRET_BLOCK": secretBlock,
		"PASSWORD_CLAIM_SIGNATURE":    signature,
		"SECRET_HASH":                 "",
	}
	return &AuthChallenge{
		challengeName:     CHALLENGE_NAME,
		timeStamp:         timeStamp,
		email:             userName,
		secretBlock:       secretBlock,
		signature:         signature,
		challengeResponse: res,
	}
}

func (a *AuthChallenge) Get(key string) (string, bool) {
	v, ok := a.challengeResponse[key]
	return v, ok
}

func (a *AuthChallenge) ChallengeResponse() map[string]string {
	new := make(map[string]string)
	for k, v := range a.challengeResponse {
		new[k] = v
	}
	return new
}

func (a *AuthChallenge) SetSecretHash(secret string) {
	a.challengeResponse["SECRET_HASH"] = secret
}

func (a *AuthChallenge) ChallengeName() string {
	return a.challengeName
}

func (a *AuthChallenge) TimeStamp() string {
	return a.timeStamp
}

func (a *AuthChallenge) Email() valueObjects.Email {
	return a.email
}

func (a *AuthChallenge) SecretBlock() string {
	return a.secretBlock
}

func (a *AuthChallenge) Signature() string {
	return a.signature
}
