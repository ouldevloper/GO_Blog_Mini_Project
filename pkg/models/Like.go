package models

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	// Id       int
	PostId int    `json:"parent_id" validate:"required"`
	Type   string `json:"type" validate:"required"`
}
