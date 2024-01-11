package models

type Auth struct {
	UserName string `jsonL:"username"`
	Password string `json:"password"`
}
