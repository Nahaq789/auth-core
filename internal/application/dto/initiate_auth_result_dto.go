package dto

type InitiateAuthResultDto struct {
	ChallengeName string
	SrpB          string
	Salt          string
	SecretBlock   string
	UserIdForSrp  string
}
