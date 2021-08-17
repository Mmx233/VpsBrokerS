package form

type DelVps struct {
	ID uint `json:"id" form:"id" binding:"required"`
}
