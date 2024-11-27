
package swagger

type AllOfTenantDtoApplicationIntegration struct {
	Id string `json:"id,omitempty"`
	Application *Object `json:"application,omitempty"`
	Integration *Object `json:"integration,omitempty"`
}
