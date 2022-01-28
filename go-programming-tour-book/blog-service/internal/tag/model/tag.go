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
