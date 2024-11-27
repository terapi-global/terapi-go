
package swagger

type ApplicationIntegrationListByApplicationIdResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfApplicationIntegrationListByApplicationIdResponseData `json:"data,omitempty"`
	Error_ *AllOfApplicationIntegrationListByApplicationIdResponseError_ `json:"error,omitempty"`
}
