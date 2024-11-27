
package swagger

type UnifiedContactResponse struct {
	Success bool `json:"success,omitempty"`
	Data *AllOfUnifiedContactResponseData `json:"data,omitempty"`
	Error_ *AllOfUnifiedContactResponseError_ `json:"error,omitempty"`
}
