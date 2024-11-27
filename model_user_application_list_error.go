
package swagger

type UserApplicationListError struct {
	Code *AllOfUserApplicationListErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfUserApplicationListErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
