package routers

import "github.com/gin-gonic/gin"

func SetRouters(e *gin.Engine) {
	LoadArticleRouters(e)
	LoadUserRouters(e)
}