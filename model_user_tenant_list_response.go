
package swagger

type UserTenantListResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfUserTenantListResponseData `json:"data,omitempty"`
	Error_ *AllOfUserTenantListResponseError_ `json:"error,omitempty"`
}
