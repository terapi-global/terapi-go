
package swagger

type CallActionResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfCallActionResponseData `json:"data,omitempty"`
	Error_ *AllOfCallActionResponseError_ `json:"error,omitempty"`
}
