
package swagger

type UserProfileStatisticsDto struct {
	AvailableIntegrations int32 `json:"availableIntegrations,omitempty"`
	AvailableApiCalls int32 `json:"availableApiCalls,omitempty"`
	ApplicationsCount int32 `json:"applicationsCount,omitempty"`
}
