package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	// 获取参数
	DB := common.GetDB()
	log.Println("start to work")
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	log.Println(name, telephone, password)
	// 数据验证
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg" : "手机号位数不正确",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg" : "密码不能少于6位",
		})
		return
	}
	if len(name) == 0 {
		name = util.RandString(10)  //如果名字不存在，则自动生成一个10位的随机字符串
	}
	// 判断手机号是否存在 （查询数据库）
	if isTelephoneExist(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg" : "用户已经存在",
		})
		return
	}
	// 创建用户
	newUser := model.User{
		Name: name,
		Telephone: telephone,
		Password: password,
	}
	DB.Create(&newUser)
	// 返回结果
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)  // 获取第一个匹配记录SELECT * FROM users WHERE name = 'jinzhu' limit 1;
	if user.ID != 0 {
		return true
	}
	return false
}

