
package swagger

type UpdateTenantConnectionError struct {
	Code *AllOfUpdateTenantConnectionErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfUpdateTenantConnectionErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
