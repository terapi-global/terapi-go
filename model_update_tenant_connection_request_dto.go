
package swagger

type UpdateTenantConnectionRequestDto struct {
	TenantId string `json:"tenantId,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	Scope string `json:"scope,omitempty"`
}
