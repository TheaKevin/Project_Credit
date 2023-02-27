package authentication

import (
	"errors"
	"log"
	"project_credit_sinarmas/backend/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type Repository interface {
	Login(req DataRequest) (string, error)
	ChangePassword(pass Password) error
	AuthenticateUser(cookie string) (models.UserTab, error)
}

type repository struct {
	db *gorm.DB
}

const SecretKey = "secret"

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Login(req DataRequest) (string, error) {
	var user models.UserTab
	res := r.db.Table("user_Tab").Where("email = ?", req.Email).First(&user)
	if res.Error != nil {
		log.Println("Get Data Error : ", res.Error)
		return "", errors.New("email atau password salah")
	}

	if user.Password != req.Password {
		return "", errors.New("email atau password salah")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *repository) ChangePassword(pass Password) error {
	var user models.UserTab
	res := r.db.Table("user_Tab").Where("email = ?", pass.Email).First(&user)
	if res.Error != nil {
		log.Println("Get Data Error : ", res.Error)
		return res.Error
	}

	if user.Password != pass.OldPassword {
		return errors.New("Password tidak sama")
	} else {
		res := r.db.Table("user_Tab").Where("email = ?", pass.Email).Update("password", pass.NewPassword)
		if res.Error != nil {
			log.Println("Update Data error : ", res.Error)
			return res.Error
		}
	}

	return nil
}

func (r *repository) AuthenticateUser(cookie string) (models.UserTab, error) {
	var user models.UserTab

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return user, err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	r.db.Table("user_Tab").Where("email = ?", claims.Issuer).First(&user)

	return user, nil
}
