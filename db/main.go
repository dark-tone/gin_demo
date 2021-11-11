package db

import (
	"fmt"
	"gin_demo/configs"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitDb() (*gorm.DB, error) {
	//dsn := "root:071311@tcp(127.0.0.1:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local"
	//db1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dbConfig := configs.DbConfig
	connArgs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Db)
	db, err := gorm.Open("mysql", connArgs)
	if err != nil {
		return new(gorm.DB), err
	}
	db.SingularTable(true)          //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会
	if gin.Mode() != gin.ReleaseMode {
		db.LogMode(true)                //打印sql语句
	}
	//开启连接池
	db.DB().SetMaxIdleConns(10)        //最大空闲连接
	db.DB().SetMaxOpenConns(100)      //最大连接数
	db.DB().SetConnMaxLifetime(30)      //最大生存时间(s)

	return db, nil
}

func GetDb() (conn *gorm.DB) {
	for {
		conn, _ = InitDb()
		if conn != nil {
			break
		}
		fmt.Println("本次未获取到mysql连接")
	}
	return conn
}