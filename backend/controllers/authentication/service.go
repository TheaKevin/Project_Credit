package authentication

import (
	"log"
	"net/http"
	"project_credit_sinarmas/backend/models"
)

type Service interface {
	Login(req DataRequest) (int, models.UserTab, error)
	ChangePassword(pass Password) (int, error)
}
type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Login(req DataRequest) (int, models.UserTab, error) {
	user, err := s.repo.Login(req)
	if err != nil {
		log.Println("Internal server error : ", err)
		return http.StatusUnauthorized, user, err
	}
	return http.StatusOK, user, nil
}

func (s *service) ChangePassword(pass Password) (int, error) {
	err := s.repo.ChangePassword(pass)
	if err != nil {
		log.Println("Internal server error : ", err)
		return http.StatusUnauthorized, err
	}
	return http.StatusOK, nil
}
