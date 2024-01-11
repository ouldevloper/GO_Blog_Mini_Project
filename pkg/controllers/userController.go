package controllers

import (
	"blog/pkg/models"
	"blog/pkg/repositories"
	"blog/pkg/utils"

	"github.com/gin-gonic/gin"
)

var uRpos = repositories.UserRepository{}

func Register(c *gin.Context) {
	var user models.User
	utils.ParseBody(c.Request, &user)
	uRpos.Register(&user)
	c.JSON(200, user)
}

func Login(c *gin.Context) {
	var auth models.Auth
	utils.ParseBody(c.Request, &auth)
	token := uRpos.Login(auth.UserName, auth.Password)
	c.JSON(200, map[string]string{"token": token})
}
