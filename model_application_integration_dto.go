
package swagger

type ApplicationIntegrationDto struct {
	Id string `json:"id,omitempty"`
	Application *AllOfApplicationIntegrationDtoApplication `json:"application,omitempty"`
	Integration *AllOfApplicationIntegrationDtoIntegration `json:"integration,omitempty"`
}
