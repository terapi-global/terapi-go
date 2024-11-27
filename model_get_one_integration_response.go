
package swagger

type GetOneIntegrationResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfGetOneIntegrationResponseData `json:"data,omitempty"`
	Error_ *AllOfGetOneIntegrationResponseError_ `json:"error,omitempty"`
}
