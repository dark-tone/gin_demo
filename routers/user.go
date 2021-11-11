package routers

import (
	"gin_demo/controllers/user"
	"github.com/gin-gonic/gin"
)

func LoadUserRouters(e *gin.Engine) {
	userGroup := e.Group("/user")

	userGroup.GET("", user.Get)
	userGroup.POST("/register", user.Register)
	userGroup.POST("/login", user.Login)
	userGroup.POST("/loginOut", user.LoginOut)
}
