package routers

import (
	"gin_demo/controllers/article"
	"gin_demo/middlewares"
	"github.com/gin-gonic/gin"
)


func LoadArticleRouters(e *gin.Engine) {
	articleGroup := e.Group("/article")

	articleGroup.GET("/:id", middlewares.CustomMiddleWare, article.Get)
	articleGroup.POST("", middlewares.LoginMiddleWare, article.Create)
	articleGroup.GET("", article.GetList)
	articleGroup.DELETE("/:id", middlewares.LoginMiddleWare, article.Delete)
}