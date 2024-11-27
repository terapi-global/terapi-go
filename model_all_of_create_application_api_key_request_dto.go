
package swagger

type AllOfCreateApplicationApiKeyRequestDto struct {
	WebUrl string `json:"webUrl,omitempty"`
	IpAddresses []string `json:"ipAddresses,omitempty"`
	ApplicationId string `json:"applicationId,omitempty"`
}
