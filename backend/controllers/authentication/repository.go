package authentication

import (
	"errors"
	"log"
	"project_credit_sinarmas/backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	Login(req DataRequest) (models.UserTab, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Login(req DataRequest) (models.UserTab, error) {
	var user models.UserTab
	res := r.db.Table("user_tab").Where("email = ?", req.Email).First(&user)
	if res.Error != nil {
		log.Println("Get Data Error : ", res.Error)
		return user, errors.New("Email atau password salah!")
	}

	if user.Password != req.Password {
		return user, errors.New("Email atau password salah!")
	}

	return user, nil
}
