
package swagger

type RevokeApplicationApiKeyResponse struct {
	Data *Object `json:"data,omitempty"`
	Success bool `json:"success,omitempty"`
	Error_ *AllOfRevokeApplicationApiKeyResponseError_ `json:"error,omitempty"`
}
