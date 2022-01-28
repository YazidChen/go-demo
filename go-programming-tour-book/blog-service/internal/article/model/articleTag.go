package model

import "yazidchen.com/go-programming-tour-book/blog-service/internal/base"

type ArticleTag struct {
	*base.Model
	ArticleID uint32 `json:"article_id"`
	TagID     uint32 `json:"tag_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
