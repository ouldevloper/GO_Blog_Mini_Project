package repositories

import (
	"blog/pkg/config"
	"blog/pkg/models"
	"crypto/sha1"
	"encoding/base64"
)

type UserRepository struct{}

func (repos *UserRepository) Register(user *models.User) {
	hasher := sha1.New()
	hasher.Write([]byte(user.Password))
	user.Password = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	config.GetCnx().Create(&user)
}

func (repos *UserRepository) login(username string, password string) {
	var user *models.User
	config.GetCnx().
		Model(models.User{}).
		Where("username=?", username).
		First(&user)
	if user == nil {
		panic("UserNot found")
	}

}

func (repos *UserRepository) ResetPassword(old string, newPass string) {

}
