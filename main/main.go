package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)
// _的作用就更特殊。当导入一个包的时候，该包的init和其他函数都会被导入；
// 但不是所有函数都需要. "_"符号可以只导入init，而不需要导入其他函数。
type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password string `gorm:"size:255;not null"`
}
func main() {
	db := InitDB()
	defer db.Close()
	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {
		// 获取参数
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
			name = RandString(10)  //如果名字不存在，则自动生成一个10位的随机字符串
		}
		// 判断手机号是否存在 （查询数据库）
		if isTelephoneExist(db, telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": 422,
				"msg" : "用户已经存在",
			})
			return
		}
		// 创建用户
		newUser := User{
			Name: name,
			Telephone: telephone,
			Password: password,
		}
		db.Create(&newUser)
		// 返回结果
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
func RandString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmMNBVCXZPOIUYTREWQASDFGHJKL")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i:= range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// 封装链接数据库的操作
func InitDB() *gorm.DB {
	driveName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "admin"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driveName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&User{})  // 自动创建数据表
	return db
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user User
	db.Where("telephone = ?", telephone).First(&user)  // 获取第一个匹配记录SELECT * FROM users WHERE name = 'jinzhu' limit 1;
	if user.ID != 0 {
		return true
	}
	return false
}