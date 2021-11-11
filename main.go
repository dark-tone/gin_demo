package main

import (
	"fmt"
	"gin_demo/configs"
	"gin_demo/routers"
	"gin_demo/validators"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
)

func main() {
	// 设置环境并加载相关配置
	gin.SetMode(gin.DebugMode)
	configs.InitConfig()

	r := gin.Default()

	// 使用redis来存储session信息
	//gob.Register(db.User{})
	redisAddress := fmt.Sprintf("%v:%v", configs.RedisConfig.Host, strconv.Itoa(configs.RedisConfig.Port))
	store, _ := redis.NewStore(10, "tcp", redisAddress, "", []byte("secret"))
	//路由上加入session中间件
	r.Use(sessions.Sessions("mySession", store))

	// 设置路由
	routers.SetRouters(r)

	// 加载自定义参数校验方法
	validators.InitValidators()

	// 初始化数据库
	//if err := db.InitDb(); err != nil {
	//	// todo：打印日志
	//}

	//if err != nil {
	//	fmt.Println("出错了")
	//} else {
	//	fmt.Println("连接成功")
	//}
	//db.SingularTable(true)
	//db.LogMode(true)
	//sqlDB := db.DB()
	//sqlDB.SetMaxIdleConns()
	//user := dbmodle.User  {
	//	Name: "Mike",
	//}
	//db.Create(&user)
	//hasThisValue := business.NewRecord(&user)
	//fmt.Println(hasThisValue)
	//business, err := gorm.Open(sqlite.Open("test.business"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//// 迁移 schema
	//business.AutoMigrate(&Product{})
	//
	//// Create
	//business.Create(&Product{Code: "D42", Price: 100})

	r.Run(":8081")
}