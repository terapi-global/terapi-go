
package swagger

type AllOfInviteTenantByApplicationIntegrationIdRequestDto struct {
	InvitedEmailAddress string `json:"invitedEmailAddress,omitempty"`
	ApplicationIntegrationId string `json:"applicationIntegrationId,omitempty"`
	IsPublicIntegration bool `json:"isPublicIntegration,omitempty"`
}
