
package swagger

type AllOfInviteTenantResponseData struct {
	ProvidedName string `json:"providedName,omitempty"`
	InvitedEmailAddress string `json:"invitedEmailAddress,omitempty"`
	//   0 = Invited  1 = Accepted  2 = Declined
	InvitationStatus *Object `json:"invitationStatus,omitempty"`
	ApplicationIntegration *Object `json:"applicationIntegration,omitempty"`
	ClientSecret string `json:"clientSecret,omitempty"`
	ClientId string `json:"clientId,omitempty"`
	WebhookUrl string `json:"webhookUrl,omitempty"`
	AuthorizedOriginUrl string `json:"authorizedOriginUrl,omitempty"`
	Id string `json:"id,omitempty"`
}
