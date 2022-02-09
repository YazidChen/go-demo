package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"yazidchen.com/go-programming-tour-book/blog-service/internal/tag/model"
	"yazidchen.com/go-programming-tour-book/blog-service/pkg/app"
	"yazidchen.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(ctx *gin.Context) {

}

func (t Tag) List(ctx *gin.Context) {
	param := model.TagGetReq{}
	resp := app.NewResponse(ctx)
	valid, vErrs := app.BindAndValid(ctx, &param)
	if !valid {
		log.Error().Err(&vErrs).Msg("app.BindAndValid errs")
		resp.ToError(errcode.InvalidParams.WithDetails(vErrs.Errors()...))
		return
	}
	resp.ToData(gin.H{})
	return
}

func (t Tag) Create(ctx *gin.Context) {

}

func (t Tag) Update(ctx *gin.Context) {

}

func (t Tag) Delete(ctx *gin.Context) {

}
