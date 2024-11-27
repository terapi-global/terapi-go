
package swagger

type AllOfUnifiedContactResponseData struct {
	ExternalId string `json:"externalId"`
	CompanyName string `json:"companyName"`
	CompanyId string `json:"companyId"`
	PlatformSource string `json:"platformSource"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	PhoneNumbers []string `json:"phoneNumbers"`
	Email string `json:"email"`
	ProfileUrl string `json:"profileUrl"`
}
