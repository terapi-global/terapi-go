
package swagger

type UpdateApplicationResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfUpdateApplicationResponseData `json:"data,omitempty"`
	Error_ *AllOfUpdateApplicationResponseError_ `json:"error,omitempty"`
}
