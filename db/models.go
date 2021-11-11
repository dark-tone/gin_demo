package db

import "time"

type DefaultDbModel struct {
	Id int `gorm:"primaryKey;autoIncrement"`
	CreateTime time.Time `gorm:"column:create_time;default:autoCreateTime"`
	UpdateTime time.Time `gorm:"default:autoUpdateTime"`
	IsDelete int `gorm:"default:0"`
}

type User struct {
	Id int `gorm:"primaryKey;autoIncrement"`
	Name string
	Avatar string
	Password string
	IsAdmin bool
	CreateTime time.Time `gorm:"column:create_time;default:autoCreateTime"`
	UpdateTime time.Time `gorm:"default:autoUpdateTime"`
}

type Article struct {
	DefaultDbModel
	Title string
	Content string
	IsDraft bool
	AuthorId int
	Category Category
}

type Category struct {
	Id int `gorm:"primaryKey;autoIncrement"`
	Name string
	AuthorId int
	CreateTime time.Time `gorm:"column:create_time;default:autoCreateTime"`
	UpdateTime time.Time `gorm:"default:autoUpdateTime"`
}
