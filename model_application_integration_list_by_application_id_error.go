
package swagger

type ApplicationIntegrationListByApplicationIdError struct {
	Code *AllOfApplicationIntegrationListByApplicationIdErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfApplicationIntegrationListByApplicationIdErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
