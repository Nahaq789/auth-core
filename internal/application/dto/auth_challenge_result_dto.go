package dto

type AuthChallengeResultDto struct {
	AccessToken  string `json:"access_token" binding:"required"`
	IdToken      string `json:"id_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}
