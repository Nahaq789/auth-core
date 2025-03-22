package dto

type AuthChallengeDto struct {
	TimeStamp   string `json:"time_stamp" binding:"required"`
	Email       string `json:"email" binding:"required"`
	SecretBlock string `json:"secret_block" binding:"required"`
	Signature   string `json:"signature" binding:"required"`
}
