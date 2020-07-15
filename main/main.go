package main

import (
	"ginEssential"
	"ginEssential/common"
	"github.com/gin-gonic/gin"
)
// _的作用就更特殊。当导入一个包的时候，该包的init和其他函数都会被导入；
// 但不是所有函数都需要. "_"符号可以只导入init，而不需要导入其他函数。

func main() {
	db := common.InitDB()  // 这里的DB是如何创建的。
	defer db.Close()
	r := gin.Default()
	r = ginEssential.CollectRoute(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}


