
package swagger

type InviteTenantError struct {
	Code *AllOfInviteTenantErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfInviteTenantErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
