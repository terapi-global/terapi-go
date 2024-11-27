
package swagger

type AddIntegrationToApplicationError struct {
	Code *AllOfAddIntegrationToApplicationErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfAddIntegrationToApplicationErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
