package v1

import (
	"github.com/gin-gonic/gin"
	"yazidchen.com/go-programming-tour-book/blog-service/pkg/app"
	"yazidchen.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(ctx *gin.Context) {
	app.NewResponse(ctx).ToError(errcode.ServerError)
	return
}

func (a Article) List(ctx *gin.Context) {

}

func (a Article) Create(ctx *gin.Context) {

}

func (a Article) Update(ctx *gin.Context) {

}

func (a Article) Delete(ctx *gin.Context) {

}
