
package swagger

type AllOfApplicationIntegrationListByApplicationIdResponseData struct {
	TotalRecords int32 `json:"totalRecords"`
	CurrentPage int32 `json:"currentPage"`
	PerPage int32 `json:"perPage"`
	Items []ApplicationIntegrationDto `json:"items"`
}
