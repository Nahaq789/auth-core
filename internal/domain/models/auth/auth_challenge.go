package auth

import valueObjects "github.com/auth-core/internal/domain/value_objects"

type AuthChallenge struct {
	challengeName     string
	timeStamp         string
	userName          valueObjects.Email
	secretBlock       string
	signature         string
	challengeResponse map[string]string
}

func NewAuthChallenge(
	challengeName string,
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
	}
	return &AuthChallenge{
		challengeName:     challengeName,
		timeStamp:         timeStamp,
		userName:          userName,
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

func (a *AuthChallenge) ChallengeName() string {
	return a.challengeName
}

func (a *AuthChallenge) TimeStamp() string {
	return a.timeStamp
}

func (a *AuthChallenge) UserName() valueObjects.Email {
	return a.userName
}

func (a *AuthChallenge) SecretBlock() string {
	return a.secretBlock
}

func (a *AuthChallenge) Signature() string {
	return a.signature
}
