
package swagger

type OAuth2ChallengeDto struct {
	RedirectUrl string `json:"redirectUrl,omitempty"`
	Tokens *AllOfOAuth2ChallengeDtoTokens `json:"tokens,omitempty"`
}
