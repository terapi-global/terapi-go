
package swagger

type GetAuthenticationTokenResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfGetAuthenticationTokenResponseData `json:"data,omitempty"`
	Error_ *AllOfGetAuthenticationTokenResponseError_ `json:"error,omitempty"`
}
