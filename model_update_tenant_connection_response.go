
package swagger

type UpdateTenantConnectionResponse struct {
	Data *Object `json:"data,omitempty"`
	Success bool `json:"success,omitempty"`
	Error_ *AllOfUpdateTenantConnectionResponseError_ `json:"error,omitempty"`
}
