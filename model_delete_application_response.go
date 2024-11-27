
package swagger

type DeleteApplicationResponse struct {
	Data *Object `json:"data,omitempty"`
	Success bool `json:"success,omitempty"`
	Error_ *AllOfDeleteApplicationResponseError_ `json:"error,omitempty"`
}
