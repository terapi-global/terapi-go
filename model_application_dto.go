
package swagger

type ApplicationDto struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Category string `json:"category,omitempty"`
	OfficialLandingUrl string `json:"officialLandingUrl,omitempty"`
	RedirectBaseUrl string `json:"redirectBaseUrl,omitempty"`
	PrivacyPolicyUrl string `json:"privacyPolicyUrl,omitempty"`
	EndUserLicenseAgreementUrl string `json:"endUserLicenseAgreementUrl,omitempty"`
	TotalIntegrations int32 `json:"totalIntegrations,omitempty"`
	TotalTenants int32 `json:"totalTenants,omitempty"`
	TotalApiCalls int32 `json:"totalApiCalls,omitempty"`
	Id string `json:"id,omitempty"`
}
