
package swagger

type UserProfileStatisticsError struct {
	Code *AllOfUserProfileStatisticsErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfUserProfileStatisticsErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
