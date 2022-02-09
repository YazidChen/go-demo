package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yazidchen.com/go-programming-tour-book/blog-service/pkg/errcode"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToData(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToPage(list interface{}, total int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:     GetPage(r.Ctx),
			PageSize: GetPageSize(r.Ctx),
			Total:    total,
		},
	})
}

func (r *Response) ToError(err *errcode.Error) {
	res := gin.H{"code": err.Code(), "msg": err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		res["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), res)
}
