
package swagger

type OAuth2ChallengeDtoResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfOAuth2ChallengeDtoResponseData `json:"data,omitempty"`
	Error_ *AllOfOAuth2ChallengeDtoResponseError_ `json:"error,omitempty"`
}
