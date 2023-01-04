package stagingCustomer

import (
	"net/http"
	"project_credit_sinarmas/backend/models"
)

type Service interface {
	GetStagingCustomer() ([]models.StagingCustomer, int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetStagingCustomer() ([]models.StagingCustomer, int, error) {

	sc, err := s.repo.GetStagingCustomer()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return sc, http.StatusOK, nil
}
