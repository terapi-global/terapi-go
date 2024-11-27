
package swagger

type CategoryListError struct {
	Code *AllOfCategoryListErrorCode `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *AllOfCategoryListErrorType_ `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
