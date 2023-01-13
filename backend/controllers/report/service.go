package report

import (
	"log"
	"net/http"
	"project_credit_sinarmas/backend/models"
)

type Service interface {
	GetReport() ([]Result, int, error)
	GetCompany() ([]models.MstCompanyTab, int, error)
	GetBranch() ([]models.BranchTab, int, error)
	GetReportFilter(q Query) ([]Result, int, error)
}
type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetReport() ([]Result, int, error) {
	transaction, err := s.repo.GetReport()
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

func (s *service) GetReportFilter(q Query) ([]Result, int, error) {
	transaction, err := s.repo.GetReportFilter(q)
	if err != nil {
		log.Println("Internal server error : ", err)
		return nil, http.StatusInternalServerError, err
	}
	return transaction, http.StatusOK, nil
}
