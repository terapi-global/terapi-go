
package swagger

type GetAuthenticationTokenError struct {
	Code *AllOfGetAuthenticationTokenErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfGetAuthenticationTokenErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
