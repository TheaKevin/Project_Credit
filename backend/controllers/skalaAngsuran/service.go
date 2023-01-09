package skalaAngsuran

import (
	"log"
	"net/http"
	"project_credit_sinarmas/backend/models"
)

type Service interface {
	GenerateSkalaAngsuran() ([]models.CustomerDataTab, int, error)
}
type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}
func (s *service) GenerateSkalaAngsuran() ([]models.CustomerDataTab, int, error) {
	sa, err := s.repo.GenerateSkalaAngsuran()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}
	return sa, http.StatusOK, nil
}
