
package swagger

type RefreshAuthenticationTokenResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfRefreshAuthenticationTokenResponseData `json:"data,omitempty"`
	Error_ *AllOfRefreshAuthenticationTokenResponseError_ `json:"error,omitempty"`
}
