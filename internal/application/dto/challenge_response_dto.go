package dto

type ChallengeResponseDto struct {
	ChallengeName string
	SrpB          string
	Salt          string
	SecretBlock   string
	UserIdForSrp  string
}
