package controller

import (
	"ginEssential/common"
	"ginEssential/model"
	"ginEssential/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
	// 2020年7月15日 加密用户的密码
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":500,
			"msg": "密码加密失败",
		})
		return
	}
	newUser := model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hashPassword),
	}
	DB.Create(&newUser)
	// 返回结果
	c.JSON(200, gin.H{
		"code":200,
		"message": "注册成功",
	})
}

func Login(c *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
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
	// 判断手机号是否存在
	user := model.User{}
	DB.Where("telephone = ?", telephone).Find(&user)
	if user.ID == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg": "手机号不存在",
		})
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":400,
			"msg":"密码错误",
		})
		return
	}
	// 发放token
	// token := "11"
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code":500, "msg":"系统异常，token生成失败"})
		log.Printf("token generate error %v", err)
		return
	}
	// 返回结果
	c.JSON(200, gin.H{
		"code": 200,
		"data":gin.H{"token":token},
		"message": "登录成功",
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

