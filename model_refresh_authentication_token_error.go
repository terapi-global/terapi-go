
package swagger

type RefreshAuthenticationTokenError struct {
	Code *AllOfRefreshAuthenticationTokenErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfRefreshAuthenticationTokenErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
