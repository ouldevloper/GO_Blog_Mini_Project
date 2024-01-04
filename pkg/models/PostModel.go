package models

import "gorm.io/gorm"

type PostModel struct {
	gorm.Model
	Id          int
	Title       string
	Description string
	Body        string
}
