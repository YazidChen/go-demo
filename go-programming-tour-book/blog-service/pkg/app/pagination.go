package app

import (
	"github.com/gin-gonic/gin"
	"yazidchen.com/go-programming-tour-book/blog-service/global"
	"yazidchen.com/go-programming-tour-book/blog-service/pkg/convert"
)

func GetPage(c *gin.Context) int {
	p := convert.StrTo(c.Query("page")).MustInt()
	if p <= 0 {
		return 1
	}
	return p
}

func GetPageSize(c *gin.Context) int {
	s := convert.StrTo(c.Query("page_size")).MustInt()
	if s <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if s > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return s
}

func GetPageOffset(page, pageSize int) int {
	o := 0
	if page > 0 {
		o = (page - 1) * pageSize
	}
	return o
}
