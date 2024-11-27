
package swagger

type InviteTenantRequestDto struct {
	InvitedEmailAddress string `json:"invitedEmailAddress,omitempty"`
	ApplicationId string `json:"applicationId,omitempty"`
	IntegrationId string `json:"integrationId,omitempty"`
	IsPublicIntegration bool `json:"isPublicIntegration,omitempty"`
}
