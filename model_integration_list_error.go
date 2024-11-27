
package swagger

type IntegrationListError struct {
	Code *AllOfIntegrationListErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfIntegrationListErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
