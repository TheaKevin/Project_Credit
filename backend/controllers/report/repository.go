package report

import (
	"log"
	"project_credit_sinarmas/backend/models"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	GetReport() ([]Result, error)
	GetCompanyData() ([]models.MstCompanyTab, error)
	GetBranchData() ([]models.BranchTab, error)
	GetReportFilter(q Query) ([]Result, error)
}

type repository struct {
	db *gorm.DB
}

type Result struct {
	PPK               string
	Name              string
	ChannelingCompany string
	DrawdownDate      time.Time
	LoanAmount        string
	LoanPeriod        string
	InterestEffective float32
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetReport() ([]Result, error) {
	var results []Result
	res := r.db.
		Table("customer_data_tab").
		Select(`
			customer_data_tab.custcode,
			customer_data_tab.ppk,
			customer_data_tab.name,
			customer_data_tab.channeling_company,
			customer_data_tab.drawdown_date,
			loan_data_tab.loan_amount,
			loan_data_tab.loan_period,
			loan_data_tab.interest_effective
		`).
		Joins("FULL OUTER JOIN loan_data_tab ON customer_data_tab.custcode = loan_data_tab.custcode").
		Where("customer_data_tab.approval_status = ? OR customer_data_tab.approval_status = ?", "0", "1").
		Scan(&results)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}
	return results, nil
}

func (r *repository) GetCompanyData() ([]models.MstCompanyTab, error) {
	var company []models.MstCompanyTab
	res := r.db.Find(&company)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}
	return company, nil
}

func (r *repository) GetBranchData() ([]models.BranchTab, error) {
	var branch []models.BranchTab
	res := r.db.Find(&branch)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}
	return branch, nil
}

func (r *repository) GetReportFilter(q Query) ([]Result, error) {
	var results []Result
	res := r.db.
		Table("customer_data_tab").
		Select(`
			customer_data_tab.custcode,
			customer_data_tab.ppk,
			customer_data_tab.name,
			customer_data_tab.channeling_company,
			customer_data_tab.drawdown_date,
			loan_data_tab.loan_amount,
			loan_data_tab.loan_period,
			loan_data_tab.interest_effective
		`).
		Joins("FULL OUTER JOIN loan_data_tab ON customer_data_tab.custcode = loan_data_tab.custcode").
		Where("customer_data_tab.approval_status = ?", q.status).
		Where("customer_data_tab.channeling_company = ?", q.company).
		Where("loan_data_tab.branch = ?", q.branch).
		Where("customer_data_tab.drawdown_date BETWEEN ? AND ?", q.start, q.end).
		Scan(&results)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}
	return results, nil
}
