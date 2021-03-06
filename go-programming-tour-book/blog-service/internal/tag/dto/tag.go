package dto

type TagGetReq struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagCreateReq struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagUpdateReq struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"required,min=3,max=100"`
	State      *uint8 `form:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type TagDelReq struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
