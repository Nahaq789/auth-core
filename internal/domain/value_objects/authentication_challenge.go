package valueObjects

import "fmt"

const (
	CHALLENGE_NAME  = "PASSWORD_VERIFIER"
	SRP_B           = "SRP_B"
	SALT            = "SALT"
	SECRET_BLOCK    = "SECRET_BLOCK"
	USER_ID_FOR_SRP = "USER_ID_FOR_SRP"
)

type AuthenticationChallenge struct {
	challengeName   string
	challengeParams map[string]string
}

func NewAuthenticationChallenge(
	challengeName string,
	challengeParams map[string]string,
) (*AuthenticationChallenge, error) {
	if !challengeValidation(challengeName) {
		return nil, fmt.Errorf("Challenge name do not match: %s", challengeName)
	}

	if err := paramsValidation(challengeParams); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &AuthenticationChallenge{
		challengeName:   challengeName,
		challengeParams: challengeParams,
	}, nil
}

func (ac *AuthenticationChallenge) GetChallengeName() string {
	return ac.challengeName
}

func (ac *AuthenticationChallenge) GetChallengeParams() map[string]string {
	return ac.challengeParams
}

func challengeValidation(name string) bool {
	return name == CHALLENGE_NAME
}

func (ac *AuthenticationChallenge) GetSrpB() string {
	return ac.challengeParams[SRP_B]
}

func (ac *AuthenticationChallenge) GetSalt() string {
	return ac.challengeParams[SALT]
}

func (ac *AuthenticationChallenge) GetSecretBlock() string {
	return ac.challengeParams[SECRET_BLOCK]
}

func (ac *AuthenticationChallenge) GetUserIdForSrp() string {
	return ac.challengeParams[USER_ID_FOR_SRP]
}

func paramsValidation(params map[string]string) error {
	requiredParams := []string{
		SRP_B,
		SALT,
		SECRET_BLOCK,
		USER_ID_FOR_SRP,
	}

	for _, param := range requiredParams {
		if value, exist := params[param]; !exist || value == "" {
			return fmt.Errorf("required parameter missing or empty: %s", param)
		}
	}

	return nil
}
