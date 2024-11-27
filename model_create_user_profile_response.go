
package swagger

type CreateUserProfileResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfCreateUserProfileResponseData `json:"data,omitempty"`
	Error_ *AllOfCreateUserProfileResponseError_ `json:"error,omitempty"`
}
