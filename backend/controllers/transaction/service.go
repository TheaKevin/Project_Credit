package transaction

import (
	"log"
	"net/http"
	"project_credit_sinarmas/backend/models"
)

type Service interface {
	GetTransaction() ([]Result, int, error)
	GetCompany() ([]models.MstCompanyTab, int, error)
	GetBranch() ([]models.BranchTab, int, error)
	GetTransactionFilter(q Query) ([]Result, int, error)
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

func (s *service) GetCompany() ([]models.MstCompanyTab, int, error) {
	company, err := s.repo.GetCompanyData()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}
	return company, http.StatusOK, nil
}

func (s *service) GetBranch() ([]models.BranchTab, int, error) {
	branch, err := s.repo.GetBranchData()
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}
	return branch, http.StatusOK, nil
}

func (s *service) GetTransactionFilter(q Query) ([]Result, int, error) {
	transaction, err := s.repo.GetTransactionFilter(q)
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
