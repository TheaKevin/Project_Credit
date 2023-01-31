package authentication

import (
	"errors"
	"log"
	"project_credit_sinarmas/backend/models"

	"gorm.io/gorm"
)

type Repository interface {
	Login(req DataRequest) (models.UserTab, error)
	ChangePassword(pass Password) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Login(req DataRequest) (models.UserTab, error) {
	var user models.UserTab
	res := r.db.Table("user_Tab").Where("email = ?", req.Email).First(&user)
	if res.Error != nil {
		log.Println("Get Data Error : ", res.Error)
		return user, errors.New("email atau password salah")
	}

	if user.Password != req.Password {
		return user, errors.New("email atau password salah")
	}

	return user, nil
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
