
package swagger

type IntegrationEndpointDto struct {
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Fields []IntegrationFieldDto `json:"fields,omitempty"`
}
