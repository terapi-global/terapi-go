
package swagger

type CategoryDto struct {
	Title string `json:"title,omitempty"`
	IsParent bool `json:"isParent,omitempty"`
	ParentId string `json:"parentId,omitempty"`
	Id string `json:"id,omitempty"`
}
