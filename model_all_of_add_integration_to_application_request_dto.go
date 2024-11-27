
package swagger

type AllOfAddIntegrationToApplicationRequestDto struct {
	ApplicationId string `json:"applicationId,omitempty"`
	IntegrationId string `json:"integrationId,omitempty"`
	IsPublicIntegration bool `json:"isPublicIntegration,omitempty"`
}