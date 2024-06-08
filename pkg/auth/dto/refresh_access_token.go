package dto

type RefreshAccessTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required" message:"auth_invalid_refresh_token"`
}

type RefreshAccessTokenResponse struct {
	AccessToken string `json:"accessToken"`
}
