
package swagger

type UserDto struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
	Username string `json:"username,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Id string `json:"id,omitempty"`
}
