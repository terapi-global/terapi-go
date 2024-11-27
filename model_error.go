
package swagger

type ModelError struct {
	Code int32 `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
