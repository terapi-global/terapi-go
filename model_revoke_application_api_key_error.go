
package swagger

type RevokeApplicationApiKeyError struct {
	Code *AllOfRevokeApplicationApiKeyErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfRevokeApplicationApiKeyErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
