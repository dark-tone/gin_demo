package models

// ResponseResult 返回给前端的固定格式
type ResponseResult struct {
	Code int
	Msg string
	Data interface{}
}

// PageModel 分页
type PageModel struct {
	PageSize int `form:"PageSize" binding:"required" default:"20"`
	PageIndex int `form:"PageIndex" binding:"required" default:"1"`
	Order int `form:"Order"`
	Asc bool `form:"Asc"`
}

type IdModel struct {
	Id int `uri:"id" form:"id" binding:"required"`
}