package models

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	Id       int
	ParentId int
	Type     string
}
