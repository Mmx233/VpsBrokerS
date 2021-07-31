package form

type VpsInit struct {
	Name string `json:"name" form:"name" binding:"required"`
}
