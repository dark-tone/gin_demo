package user

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"gin_demo/db"
	"gin_demo/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInfo struct {
	User string `form:"User" binding:"required"`
	Password string `form:"Password" binding:"required"`
}

type RegisterInfo struct {
	User string `form:"User" binding:"required"`
	Password string `form:"Password" binding:"required"`
	Avatar string `form:"Avatar" binding:"required"`
}

func Get(ctx *gin.Context) {
	session :=  sessions.Default(ctx)
	userSession := session.Get("userInfo")
	userJson, ok := userSession.([]byte)
	if !ok {
		ctx.JSON(http.StatusForbidden, models.ResponseResult{Code: -1, Msg: "找不到对应的用户信息"})
		return
	}
	var user db.User
	err := json.Unmarshal(userJson, &user)
	if err != nil {
		ctx.JSON(http.StatusForbidden, models.ResponseResult{Code: -1, Msg: "找不到对应的用户信息"})
		return
	}
	ctx.JSON(http.StatusOK, models.ResponseResult{Data: user})
}

// Login 登录
func Login(ctx *gin.Context) {
	var user db.User
	var query LoginInfo
	if err := ctx.ShouldBind(&query); err != nil {
		ctx.JSON(http.StatusForbidden, models.ResponseResult{Code: -1, Msg: err.Error()})
		return
	}
	result := db.GetDb().Limit(1).Where("name = ?", query.User).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusForbidden, models.ResponseResult{Code: -1, Msg: result.Error.Error()})
		return
	}
	// 获取密码的md5
	data := []byte(query.Password)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	if user.Password != md5str {
		ctx.JSON(http.StatusForbidden, models.ResponseResult{Code: -1, Msg: "用户名或密码错误"})
		return
	}
	// 设置用户信息至session
	session := sessions.Default(ctx)
	userJson, _ := json.Marshal(user)
	session.Set("userInfo", userJson)
	session.Save()
	// 这里登录成功跳转到主页（暂时没有）
	ctx.Redirect(http.StatusFound, "/")
}

// LoginOut 退出登录
func LoginOut(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	ctx.Redirect(http.StatusFound, "/")
}

// Register 注册
func Register(ctx *gin.Context) {
	var user db.User
	var query RegisterInfo
	if err := ctx.ShouldBind(&query); err != nil {
		ctx.JSON(http.StatusForbidden, models.ResponseResult{Code: -1, Msg: err.Error()})
		return
	}
	db.GetDb().Limit(1).Where("name = ?", query.User).Find(&user)
	if user.Id > 0 {
		ctx.JSON(http.StatusForbidden, models.ResponseResult{Code: -1, Msg: "已存在相同用户名"})
		return
	}
	// 获取密码的md5
	data := []byte(query.Password)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)

	user.Name = query.User
	user.Password = md5str
	user.Avatar = query.Avatar
	db.GetDb().Create(&user)

	// 设置用户信息至session
	session := sessions.Default(ctx)
	session.Set("userInfo", user)
	session.Save()
	ctx.Redirect(http.StatusFound, "/")
}