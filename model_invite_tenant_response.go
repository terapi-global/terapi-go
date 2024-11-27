
package swagger

type InviteTenantResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfInviteTenantResponseData `json:"data,omitempty"`
	Error_ *AllOfInviteTenantResponseError_ `json:"error,omitempty"`
}
