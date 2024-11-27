
package swagger

type TenantDto struct {
	ProvidedName string `json:"providedName,omitempty"`
	InvitedEmailAddress string `json:"invitedEmailAddress,omitempty"`
	//   0 = Invited  1 = Accepted  2 = Declined
	InvitationStatus *AllOfTenantDtoInvitationStatus `json:"invitationStatus,omitempty"`
	ApplicationIntegration *AllOfTenantDtoApplicationIntegration `json:"applicationIntegration,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	WebhookUrl string `json:"webhookUrl,omitempty"`
	AuthorizedOriginUrl string `json:"authorizedOriginUrl,omitempty"`
	Id string `json:"id,omitempty"`
}
