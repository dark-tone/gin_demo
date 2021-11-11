package article

import (
	"encoding/json"
	"gin_demo/db"
	"gin_demo/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type articleCreateInfo struct {
	Title string `form:"Title" binding:"required,max=15"`
	Content string `form:"Content" binding:"required"`
}

type articleGetListInfo struct {
	AuthorId int `form:"AuthorId" binding:"required"`
	models.PageModel
}

func Get(ctx *gin.Context) {
	var article db.Article
	var query models.IdModel
	if err := ctx.ShouldBindUri(&query); err != nil {
		ctx.JSON(http.StatusForbidden, models.ResponseResult{Code: -1, Msg: err.Error()})
		return
	}
	db.GetDb().Limit(1).Find(&article, query.Id)
	ctx.JSON(http.StatusOK, models.ResponseResult{Data: article})
}

func Create(ctx *gin.Context) {
	var articleInfo articleCreateInfo
	//ctx.Bind(&article)
	if err := ctx.Bind(&articleInfo); err != nil {
		ctx.JSON(http.StatusForbidden, models.ResponseResult{Code: -1, Msg: err.Error()})
		return
	}
	article := db.Article{
		Title: articleInfo.Title,
		Content: articleInfo.Content,
	}
	var user db.User
	session := sessions.Default(ctx)
	userJson := session.Get("userInfo").([]byte)
	json.Unmarshal(userJson, &user)
	article.AuthorId = user.Id
	db.GetDb().Create(&article)
	ctx.JSON(http.StatusOK, models.ResponseResult{})
}

// Delete 逻辑删除
func Delete(ctx *gin.Context) {
	var query models.IdModel
	if err := ctx.ShouldBind(&query); err != nil {
		ctx.JSON(http.StatusForbidden, models.ResponseResult{Code: -1, Msg: err.Error()})
		return
	}
	db.GetDb().Model(&db.Article{}).Where(query.Id).UpdateColumn("is_delete", true)
	ctx.JSON(http.StatusOK, models.ResponseResult{})
}

// GetList 获取整个分页列表
//func GetList(ctx *gin.Context) {
//	var articles []db.Article
//	var query articleGetListInfo
//	if err := ctx.ShouldBind(&query); err != nil {
//		ctx.JSON(http.StatusForbidden, models.ResponseResult{Code: -1, Msg: err.Error()})
//		return
//	}
//	result := db.GetDb().
//		Limit(query.PageSize).
//		Offset((query.PageIndex - 1) * query.PageIndex).
//		Where("author_id = ? and is_draft = false", query.AuthorId).
//		Find(&articles)
//
//	if result.Error != nil {
//		ctx.JSON(http.StatusForbidden, models.ResponseResult{Code: -1, Msg: result.Error.Error()})
//		return
//	}
//	ctx.JSON(http.StatusOK, models.ResponseResult{Data: articles})
//}

type FFF struct {
	db.Article
	db.Category
}

func GetList(ctx *gin.Context) {
	var articles []db.Article
	var query articleGetListInfo
	if err := ctx.ShouldBind(&query); err != nil {
		ctx.JSON(http.StatusForbidden, models.ResponseResult{Code: -1, Msg: err.Error()})
		return
	}
	result := db.GetDb().
		Limit(query.PageSize).
		Offset((query.PageIndex - 1) * query.PageIndex).
		Model(&db.Article{}).
		Where("author_id = ?", query.AuthorId).Find(&articles)
		//Where("article.author_id = ? and article.is_draft = false", query.AuthorId).
		//Joins("left join category on article.category_id = category.id").
		//Scan(&articles)

	if result.Error != nil {
		ctx.JSON(http.StatusForbidden, models.ResponseResult{Code: -1, Msg: result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, models.ResponseResult{Data: articles})
}