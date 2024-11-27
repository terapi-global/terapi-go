
package swagger

type CreateApplicationRequestDto struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Category string `json:"category,omitempty"`
	OfficialLandingUrl string `json:"officialLandingUrl,omitempty"`
	RedirectBaseUrl string `json:"redirectBaseUrl,omitempty"`
	PrivacyPolicyUrl string `json:"privacyPolicyUrl,omitempty"`
	EndUserLicenseAgreementUrl string `json:"endUserLicenseAgreementUrl,omitempty"`
}
