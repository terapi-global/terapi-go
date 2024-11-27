
package swagger
import (
	"time"
)

type AllOfApplicationIntegrationDtoIntegration struct {
	Id string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	NameIdentifier string `json:"nameIdentifier,omitempty"`
	ShortDescription string `json:"shortDescription,omitempty"`
	Icon string `json:"icon,omitempty"`
	//   0 = Released  1 = Pending  2 = NotPlanned
	ReleaseStatus *Object `json:"releaseStatus,omitempty"`
	LaunchDate time.Time `json:"launchDate,omitempty"`
	DocumentationUrl string `json:"documentationUrl,omitempty"`
	SwaggerUrl string `json:"swaggerUrl,omitempty"`
	IntegrationEndpoints []IntegrationEndpointDto `json:"integrationEndpoints,omitempty"`
	IntegrationEvents []IntegrationEventDto `json:"integrationEvents,omitempty"`
}
