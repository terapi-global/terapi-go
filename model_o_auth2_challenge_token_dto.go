
package swagger

type OAuth2ChallengeTokenDto struct {
	AccessToken string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	ExpiresIn int32 `json:"expiresIn,omitempty"`
}
