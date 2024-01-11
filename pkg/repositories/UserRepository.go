package repositories

import (
	"blog/pkg/config"
	"blog/pkg/models"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your_secret_key")

type UserRepository struct{}

func crypt(s string) string {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func generateToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expiration time (1 hour in this example)
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return ""
	}

	return tokenString
}

func (repos *UserRepository) Register(user *models.User) {
	user.Password = crypt(user.Password)
	config.GetCnx().Create(&user)
}

func (repos *UserRepository) Login(username string, password string) string {
	var user *models.User
	config.GetCnx().
		Model(models.User{}).
		Where("username=?", username).
		First(&user)
	if user == nil {
		panic("User Not found")
	}

	if user.Password != crypt(password) {
		panic("Invalid Password")
	}
	return generateToken(user.Username)
}

func (repos *UserRepository) ResetPassword(old string, newPass string) {

}
