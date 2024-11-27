
package swagger

type InviteTenantByApplicationIntegrationIdError struct {
	Code *AllOfInviteTenantByApplicationIntegrationIdErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfInviteTenantByApplicationIntegrationIdErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
