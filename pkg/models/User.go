package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int    `json:"id"`
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Password string `json:"password"`
}
