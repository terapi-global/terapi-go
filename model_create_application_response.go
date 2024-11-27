
package swagger

type CreateApplicationResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfCreateApplicationResponseData `json:"data,omitempty"`
	Error_ *AllOfCreateApplicationResponseError_ `json:"error,omitempty"`
}
