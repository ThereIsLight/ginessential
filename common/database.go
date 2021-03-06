package common

import (
	"fmt"
	"ginEssential/model"
	_ "ginEssential/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB
func InitDB() *gorm.DB {
	//driveName := "mysql"
	//host := "localhost"
	//port := "3306"
	//database := "ginessential"
	//username := "root"
	//password := "admin"
	//charset := "utf8"
	driverName := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})  // 自动创建数据表

	DB = db
	return db
}
func GetDB() *gorm.DB {
	return DB
}