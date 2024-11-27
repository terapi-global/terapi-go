
package swagger

type UnifiedContactsListResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfUnifiedContactsListResponseData `json:"data,omitempty"`
	Error_ *AllOfUnifiedContactsListResponseError_ `json:"error,omitempty"`
}
