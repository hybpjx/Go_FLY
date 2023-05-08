package cmd

import (
	"fmt"
	"gofly/conf"
	"gofly/global"
	"gofly/router"
	"gofly/utils"
)

func Start() {
	var initError error

	//============== 初始化系统配置 ==============
	conf.InitConfig()

	//============== 初始化日志组件 ==============
	// 需要找个值去接受这个logger对象 不然后面调用不太好调
	global.Logger = conf.InitLogger()

	//============== 初始化数据库链接 ==============
	db, err := conf.InitDB()
	global.DB = db
	{
		// 判断初始化中是否有错误
		if err != nil {
			initError = utils.AppendError(initError, err)
		}

		if initError != nil {
			if global.Logger != nil {
				global.Logger.Error(initError.Error())
			}
			panic(initError.Error())
		}
	}

	//============== 初始化Redis连接 ==============
	redisClient, err := conf.InitRedis()
	global.RedisClient = redisClient
	{
		// 判断初始化中是否有错误
		if err != nil {
			if err != nil {
				initError = utils.AppendError(initError, err)
			}
		}

	}
	//============== 初始化系统路由 ==============
	router.InitRouter()
}

func Clean() {
	fmt.Println("++++++++++clean+++++++++")
}
