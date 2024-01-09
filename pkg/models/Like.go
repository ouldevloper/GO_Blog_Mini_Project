package models

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	ID     int    `json:"id"`
	PostId int    `json:"parent_id" validate:"required"`
	Type   string `json:"type" validate:"required"`
	User   *User  `json:"user"`
}
