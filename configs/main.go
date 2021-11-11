package configs

import "github.com/gin-gonic/gin"

var DbConfig DbConfigObj
var RedisConfig RedisObjConfigObj

func InitConfig() {
	switch gin.Mode() {
	case gin.ReleaseMode:
		// 这里放处理正式服环境情况的代码
	default:
	}
	initDbConfig()
	initRedisConfig()
}

/**
 * @description 初始化数据库配置
 */
func initDbConfig() {
	// 默认加载dev环境的配置
	DbConfig = DbConfigObj{
		Host: "localhost",
		User: "root",
		Password: "071311",
		Port: 3306,
		Db: "blog_db",
	}

	// 加载正式服环境配置
	if gin.Mode() == gin.ReleaseMode {
		DbConfig.Host = "***"
		DbConfig.User = "***"
	}
}

/**
 * @description 初始化Redis配置
 */
func initRedisConfig() {
	RedisConfig = RedisObjConfigObj{
		Host: "localhost",
		User: "",
		Password: "",
		Port: 6379,
	}

	// 加载正式服环境配置
	if gin.Mode() == gin.ReleaseMode {
		RedisConfig.Host = "localhost"
		RedisConfig.User = "root"
	}
}