
package swagger

type IntegrationListResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfIntegrationListResponseData `json:"data,omitempty"`
	Error_ *AllOfIntegrationListResponseError_ `json:"error,omitempty"`
}
