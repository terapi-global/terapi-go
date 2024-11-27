
package swagger

type UserApplicationListResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfUserApplicationListResponseData `json:"data,omitempty"`
	Error_ *AllOfUserApplicationListResponseError_ `json:"error,omitempty"`
}
