
package swagger

type CreateApplicationError struct {
	Code *AllOfCreateApplicationErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfCreateApplicationErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
