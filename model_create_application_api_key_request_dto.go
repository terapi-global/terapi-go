
package swagger

type CreateApplicationApiKeyRequestDto struct {
	WebUrl string `json:"webUrl,omitempty"`
	IpAddresses []string `json:"ipAddresses,omitempty"`
	ApplicationId string `json:"applicationId,omitempty"`
}
