package stagingCustomer

import (
	"fmt"
	"project_credit_sinarmas/backend/models"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	GetStagingCustomer() ([]models.StagingCustomer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetStagingCustomer() ([]models.StagingCustomer, error) {
	var sc []models.StagingCustomer
	var c []models.CustomerDataTab
	var com []models.MstCompanyTab
	var b []models.BranchTab
	var v []models.VehicleDataTab
	t := time.Now()

	res := r.db.Where("sc_flag = ? AND sc_create_date <= ?", "0", t.Format("2006-01-02T15:04:05Z")).Find(&sc)
	if res.Error != nil {
		return nil, res.Error
	}

	for _, item := range sc {
		res2 := r.db.Where("ppk = ?", item.CustomerPpk).First(&c)
		if res2.Error == nil {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " ppk duplicate")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		res3 := r.db.Where("company_short_name = ?", item.ScCompany).First(&com)
		if res3.Error != nil {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " sc_company gk ad")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		res4 := r.db.Where("code = ?", item.ScBranchCode).First(&b)
		if res4.Error != nil {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " branch code gk ad")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		// t2, err := time.Parse("January 2006", item.LoanTglPk)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// if t.Format("January 2006") != t2.Format("January 2006") {
		// 	fmt.Println(strconv.FormatInt(item.ID, 10) + " bulan pk berbeda dengan bulan berjalan")
		// 	r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
		// 	continue
		// }

		if item.CustomerIDType == "1" && item.CustomerIDNumber == "" {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " customer id number gk ad")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		res5 := r.db.Where("engine_no = ? OR chasis_no = ?", item.VehicleEngineNo, item.VehicleChasisNo).First(&v)
		if res5.Error == nil {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " engine atau chasis no duplikasi")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}
	}

	return sc, nil
}
