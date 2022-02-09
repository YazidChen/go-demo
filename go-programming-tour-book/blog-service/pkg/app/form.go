package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key string
	Msg string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Msg
}

func (vs *ValidErrors) Error() string {
	return strings.Join(vs.Errors(), ",")
}

func (vs ValidErrors) Errors() []string {
	var errs []string
	for _, v := range vs {
		errs = append(errs, v.Error())
	}
	return errs
}

func BindAndValid(ctx *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	// 绑定并校验，返回错误
	err := ctx.ShouldBind(v)
	if err != nil {
		t := ctx.Value("trans")
		// 接口类型断言 Type assertions
		trans, _ := t.(ut.Translator)
		vErrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}
		//对错误消息体进行翻译
		for key, value := range vErrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key: key,
				Msg: value,
			})
		}
		return false, errs
	}
	return true, nil
}
