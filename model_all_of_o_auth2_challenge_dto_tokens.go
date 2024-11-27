
package swagger

type AllOfOAuth2ChallengeDtoTokens struct {
	AccessToken string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	ExpiresIn int32 `json:"expiresIn,omitempty"`
}
