
package swagger

type AddIntegrationToApplicationResponse struct {
	Data *Object `json:"data,omitempty"`
	Success bool `json:"success,omitempty"`
	Error_ *AllOfAddIntegrationToApplicationResponseError_ `json:"error,omitempty"`
}
