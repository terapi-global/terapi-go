
package swagger

type CreateApplicationApiKeyError struct {
	Code *AllOfCreateApplicationApiKeyErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfCreateApplicationApiKeyErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
