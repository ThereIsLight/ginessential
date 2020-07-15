package middleware

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取authorization header
		tokenString := c.GetHeader("Authorization")  //从请求body中获取tokenstring
		log.Printf("token:=  %v", tokenString)
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			// 后面就不修改了。
			response.Response(c, http.StatusUnauthorized, 401, nil, "权限不足,token字符串为空")
			// c.JSON(http.StatusUnauthorized, gin.H{"code":401, "msg":"权限不足,token字符串为空"})
			c.Abort()  // 作用？？
			return
		}
		// 取出从"Bearer "之后的字符串，这才是真正的token
		tokenString = tokenString[7:]
		// 解析token
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid{
			c.JSON(http.StatusUnauthorized, gin.H{"code":401, "msg":"权限不足，token解析出错"})
			c.Abort()  // 作用？？
			return
		}
		// 通过验证
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)  // SELECT * FROM users WHERE id = userID
		log.Printf("用户ID %v", userId)
		log.Printf("用户信息 %v", user)

		// 用户不存在
		if userId == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code":401, "msg":"权限不足，用户不存在"})
			c.Abort()  // 作用？？
			return
		}
		c.Set("user", user)  //写入上下文是什么意思？？就是写入到*gin.Context中吗？
		c.Next()
	}
}
