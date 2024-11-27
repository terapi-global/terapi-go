
package swagger

type CreateUserProfileError struct {
	Code *AllOfCreateUserProfileErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfCreateUserProfileErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
