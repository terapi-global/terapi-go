
package swagger

type AllOfApplicationIntegrationListByApplicationIdResponseError_ struct {
	Code *Object `json:"code,omitempty"`
	//   1 = BusinessLogic  2 = InternalServerError
	Type_ *Object `json:"type,omitempty"`
	Values map[string]string `json:"values,omitempty"`
}
