package model

import "yazidchen.com/go-programming-tour-book/blog-service/internal/base"

type Tag struct {
	*base.Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}

type TagGetReq struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagCreateReq struct {
	Name     string `form:"name" binding:"required,min=3,max=100"`
	CreateBy string `form:"create_by" binding:"required,min=3,max=100"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagUpdateReq struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"required,min=3,max=100"`
	State      string `form:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type TagDelReq struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
