
package swagger

type UserProfileStatisticsResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfUserProfileStatisticsResponseData `json:"data,omitempty"`
	Error_ *AllOfUserProfileStatisticsResponseError_ `json:"error,omitempty"`
}
