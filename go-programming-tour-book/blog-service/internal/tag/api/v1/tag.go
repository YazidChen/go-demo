package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"yazidchen.com/go-programming-tour-book/blog-service/internal/tag/dto"
	"yazidchen.com/go-programming-tour-book/blog-service/internal/tag/service"
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
	param := dto.TagGetReq{}
	resp := app.NewResponse(ctx)
	valid, vErrs := app.BindAndValid(ctx, &param)
	if !valid {
		log.Error().Err(&vErrs).Msg("app.BindAndValid errs")
		resp.ToError(errcode.InvalidParams.WithDetails(vErrs.Errors()...))
		return
	}
	svc := service.NewTagService(ctx.Request.Context())
	pager := app.Pager{Page: app.GetPage(ctx), PageSize: app.GetPageSize(ctx)}

	list, err := svc.GetTagList(&param, &pager)
	if err != nil {
		log.Error().Err(err).Msg("TagService.GetTagList err")
		resp.ToError(errcode.ErrorGetTagListFail)
		return
	}

	total, err := svc.CountTag(&param)
	if err != nil {
		log.Error().Err(err).Msg("TagService.CountTag err")
		resp.ToError(errcode.ErrorCountTagFail)
		return
	}

	resp.ToPage(list, total)
	return
}

func (t Tag) Create(ctx *gin.Context) {
	param := dto.TagCreateReq{}
	resp := app.NewResponse(ctx)
	valid, vErrs := app.BindAndValid(ctx, &param)
	if !valid {
		log.Error().Err(&vErrs).Msg("app.BindAndValid errs")
		resp.ToError(errcode.InvalidParams.WithDetails(vErrs.Errors()...))
		return
	}
	svc := service.NewTagService(ctx.Request.Context())

	err := svc.CreateTag(&param)
	if err != nil {
		log.Error().Err(err).Msg("TagService.Create err")
		resp.ToError(errcode.ErrorCreateTagFail)
		return
	}

	resp.ToData(gin.H{})
	return
}

func (t Tag) Update(ctx *gin.Context) {

}

func (t Tag) Delete(ctx *gin.Context) {

}
