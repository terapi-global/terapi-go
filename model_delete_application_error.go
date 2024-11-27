
package swagger

type DeleteApplicationError struct {
	Code *AllOfDeleteApplicationErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfDeleteApplicationErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
