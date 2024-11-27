
package swagger

type InviteTenantByApplicationIntegrationIdResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfInviteTenantByApplicationIntegrationIdResponseData `json:"data,omitempty"`
	Error_ *AllOfInviteTenantByApplicationIntegrationIdResponseError_ `json:"error,omitempty"`
}
