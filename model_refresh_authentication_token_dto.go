
package swagger

type RefreshAuthenticationTokenDto struct {
	AccessToken string `json:"accessToken,omitempty"`
	ExpiresIn int32 `json:"expiresIn,omitempty"`
}
