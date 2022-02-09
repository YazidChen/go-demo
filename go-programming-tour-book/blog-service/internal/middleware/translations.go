package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// go-playground/universal-translator 通用翻译器
		// go-playground/validator/v10/translations validator翻译器
		uni := ut.New(en.New(), zh.New())
		locale := ctx.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zhTranslations.RegisterDefaultTranslations(v, trans)
			case "en":
				_ = enTranslations.RegisterDefaultTranslations(v, trans)
			default:
				_ = zhTranslations.RegisterDefaultTranslations(v, trans)
			}
			// 上下文存入Translator，方便后续使用
			ctx.Set("trans", trans)
		}
		ctx.Next()
	}
}
