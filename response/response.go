package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 统一封装http返回

func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{
		"code":code,
		"data":data,
		"msg":msg,
	})
	// httpStatus与code的区别：一个是标准Http状态码，一个是业务状态码。
}

func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}

func Fail(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 400, data, msg)
}
