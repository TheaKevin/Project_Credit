package transaction

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	GetTransaction() ([]Result, error)
}

type repository struct {
	db *gorm.DB
}

type Result struct {
	Custcode          string
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
func (r *repository) GetTransaction() ([]Result, error) {
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
		Where("customer_data_tab.approval_status = ?", "9").
		Scan(&results)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}
	return results, nil
}
