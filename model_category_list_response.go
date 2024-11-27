
package swagger

type CategoryListResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfCategoryListResponseData `json:"data,omitempty"`
	Error_ *AllOfCategoryListResponseError_ `json:"error,omitempty"`
}
