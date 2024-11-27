
package swagger

type GetOneIntegrationError struct {
	Code *AllOfGetOneIntegrationErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfGetOneIntegrationErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
