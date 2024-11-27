
package swagger

type Response struct {
	Data *Object `json:"data,omitempty"`
	Success bool `json:"success,omitempty"`
	Error_ *AllOfResponseError_ `json:"error,omitempty"`
}
