
package swagger

type DeclineTenantInvitationError struct {
	Code *AllOfDeclineTenantInvitationErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfDeclineTenantInvitationErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
