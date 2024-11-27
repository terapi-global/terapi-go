
package swagger

type ApplicationApiKeyByApplicationIdError struct {
	Code *AllOfApplicationApiKeyByApplicationIdErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfApplicationApiKeyByApplicationIdErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
