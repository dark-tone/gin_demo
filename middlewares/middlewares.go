package middlewares

import (
	"fmt"
	"gin_demo/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CustomMiddleWare(ctx *gin.Context) {
	ctx.Set("customKey", "gt")
	start := time.Now()
	fmt.Println("start time:" + start.String())

	ctx.Next()

	end := time.Now()
	total := time.Since(start)
	fmt.Println("end time:" + end.String())
	fmt.Println("total time:" + total.String())
}

func LoginMiddleWare(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get("userInfo")
	if user == nil {
		ctx.Abort()
		ctx.JSON(http.StatusUnauthorized, models.ResponseResult{Code: -2, Data: "未登录"})
	}
}