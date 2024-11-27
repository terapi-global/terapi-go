
package swagger

type ApplicationApiKeyByApplicationIdResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfApplicationApiKeyByApplicationIdResponseData `json:"data,omitempty"`
	Error_ *AllOfApplicationApiKeyByApplicationIdResponseError_ `json:"error,omitempty"`
}
