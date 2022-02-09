package routers

import (
	"github.com/gin-gonic/gin"
	articleV1 "yazidchen.com/go-programming-tour-book/blog-service/internal/article/api/v1"
	"yazidchen.com/go-programming-tour-book/blog-service/internal/middleware"
	tagV1 "yazidchen.com/go-programming-tour-book/blog-service/internal/tag/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())

	tag := tagV1.NewTag()
	article := articleV1.NewArticle()

	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("article/:id", article.Get)
		apiV1.GET("/articles", article.List)
	}
	return r
}
