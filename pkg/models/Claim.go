package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	// gorm.Model
	Username string `json:"username"`
	*jwt.RegisteredClaims
}
