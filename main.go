package main

import (
	"ginEssential/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)
// _的作用就更特殊。当导入一个包的时候，该包的init和其他函数都会被导入；
// 但不是所有函数都需要. "_"符号可以只导入init，而不需要导入其他函数。

func main() {
	InitConfig()  // 读取配置
	db := common.InitDB()  // 这里的DB是如何创建的。
	defer db.Close()
	r := gin.Default()
	// r = CollectRoute(r)
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	r.Run() // 如果从配置文件中读取不到端口号，默认端口号为8080
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}


