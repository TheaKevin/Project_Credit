package authentication

import (
	"errors"
	"log"
	"project_credit_sinarmas/backend/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
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
		return "", errors.New("email tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password)); err != nil {
		return "", errors.New("password salah")
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

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(pass.OldPassword)); err != nil {
		return errors.New("Password tidak sama")
	} else {
		passwordBaru, _ := bcrypt.GenerateFromPassword([]byte(pass.NewPassword), 14)
		res := r.db.Table("user_Tab").Where("email = ?", pass.Email).Update("password", passwordBaru)
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
