package transaction

import (
	"log"
	"net/http"
)

type Service interface {
	GetTransaction() ([]Result, int, error)
	UpdateTransaction(req []DataRequest) (int, error)
}
type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}
func (s *service) GetTransaction() ([]Result, int, error) {
	transaction, err := s.repo.GetTransaction()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}
	return transaction, http.StatusOK, nil
}

func (s *service) UpdateTransaction(req []DataRequest) (int, error) {
	err := s.repo.UpdateTransaction(req)
	if err != nil {
		log.Println("Internal server error : ", err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
