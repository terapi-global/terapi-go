
package swagger

type UpdateApplicationError struct {
	Code *AllOfUpdateApplicationErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfUpdateApplicationErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
