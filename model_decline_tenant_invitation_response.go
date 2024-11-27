
package swagger

type DeclineTenantInvitationResponse struct {
	Data *Object `json:"data,omitempty"`
	Success bool `json:"success,omitempty"`
	Error_ *AllOfDeclineTenantInvitationResponseError_ `json:"error,omitempty"`
}
