
package swagger

type CreateApplicationApiKeyResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfCreateApplicationApiKeyResponseData `json:"data,omitempty"`
	Error_ *AllOfCreateApplicationApiKeyResponseError_ `json:"error,omitempty"`
}
