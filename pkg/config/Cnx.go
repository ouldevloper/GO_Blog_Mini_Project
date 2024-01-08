package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

func connect() {
	d, err := gorm.Open(mysql.Open("blog_api:user@tcp(127.0.0.1:3306)/blog_api?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		// log.Fatal("DB Error : ", err.Error())
		panic(err)
	}
	db = d
}

func GetCnx() *gorm.DB {
	if db == nil {
		connect()
	}
	return db
}
