package common

import (
	"fmt"

	"github.com/Yingzhixin/gin_demo/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB 数据库实例
var DB *gorm.DB

//InitDB 初始化mysql数据库连接
func InitDB() *gorm.DB {
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database,err:" + err.Error())
	}
	db.AutoMigrate(&models.User{})
	DB = db

	return db
}

//GetDB 返回DB实例
func GetDB() *gorm.DB {
	return DB
}
