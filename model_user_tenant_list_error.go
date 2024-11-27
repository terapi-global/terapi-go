
package swagger

type UserTenantListError struct {
	Code *AllOfUserTenantListErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfUserTenantListErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
