package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	// 获取当前目录
	workDir, _ := os.Getwd()
	// 文件名
	viper.SetConfigName("settings")
	// 文件后缀
	viper.SetConfigType("yml")
	// 拼接完整目录
	viper.AddConfigPath(workDir + "/conf")
	// 读取
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("load config error: %s \n", err.Error()))
	}
	fmt.Sprintln("port is ", viper.GetString("server.port"))
}
