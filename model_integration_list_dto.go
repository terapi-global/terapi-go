
package swagger

type IntegrationListDto struct {
	TotalRecords int32 `json:"totalRecords"`
	CurrentPage int32 `json:"currentPage"`
	PerPage int32 `json:"perPage"`
	Items []IntegrationDto `json:"items"`
}
