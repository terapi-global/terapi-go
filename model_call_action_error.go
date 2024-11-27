
package swagger

type CallActionError struct {
	Code *AllOfCallActionErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfCallActionErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
