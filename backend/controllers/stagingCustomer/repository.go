package stagingCustomer

import (
	"fmt"
	"project_credit_sinarmas/backend/models"
	"strconv"
	"strings"
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
	var acc models.AppCustCodeTab
	t := time.Now()

	res := r.db.Where("sc_flag = ? AND sc_create_date <= ?", "0", t.Format("2006-01-02T15:04:05Z")).Find(&sc)
	if res.Error != nil {
		return nil, res.Error
	}

	for _, item := range sc {
		// validasi ppk duplikasi
		res2 := r.db.Where("ppk = ?", item.CustomerPpk).First(&c)
		if res2.Error == nil {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " ppk duplicate")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		// validasi SC_Company terdaftar
		res2 = r.db.Where("company_short_name = ?", item.ScCompany).First(&com)
		if res2.Error != nil {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " sc_company gk ad")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		// validasi SC_BranchCode terdaftar
		res2 = r.db.Where("code = ?", item.ScBranchCode).First(&b)
		if res2.Error != nil {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " branch code gk ad")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		// validasi TGL_PK di bulan yang sama
		date, err := time.Parse("2006-01-02", item.LoanTglPk)
		if err != nil {
			fmt.Println(err)
		}
		if (int(t.Month()) != int(date.Month())) || (int(t.Year()) != int(date.Year())) {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " bulan pk berbeda dengan bulan berjalan")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		// validasi CustomerIDType & CustomerIDNumber
		if item.CustomerIDType == "1" && item.CustomerIDNumber == "" {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " customer id number gk ad")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		// validasi Nama Debitur no char spesial
		if strings.ContainsAny(item.CustomerName, "!@#$%^&*()_+-=[]{}|\\:;'\"<>,./?") {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " nama customer terdapat special character")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		// validasi Vehicle_BPKB kosong
		if item.VehicleBpkb == "" {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " vehicle_bpkb gk ad")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		// validasi Vehicle_STNK kosong
		if item.VehicleStnk == "" {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " vehicle_stnk gk ad")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		// validasi Vehicle_Engine_NO kosong
		if item.VehicleEngineNo == "" {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " vehicle_engine_no gk ad")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		// validasi Vehicle_Chasis_NO kosong
		if item.VehicleChasisNo == "" {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " vehicle_chasis_no gk ad")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		// validasi Vehicle_Engine_NO Duplikasi
		res2 = r.db.Where("engine_no = ?", item.VehicleEngineNo).First(&v)
		if res2.Error == nil {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " engine no duplikasi")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		// validasi Vehicle_Chasis_No Duplikasi
		res2 = r.db.Where("chasis_no = ?", item.VehicleChasisNo).First(&v)
		if res2.Error == nil {
			fmt.Println(strconv.FormatInt(item.ID, 10) + " chasis no duplikasi")
			r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "8")
			continue
		}

		//Generate CustCode
		res2 = r.db.First(&acc)
		if res2.Error != nil {
			return nil, res.Error
		}
		AppCustCode := "006"
		r.db.Model(&acc).Where("id = ?", acc.ID).Update("value", acc.Value+1)
		appCustCodeSeq := "0000000000" + strconv.Itoa(acc.Value)
		appCustCodeSeq = appCustCodeSeq[len(appCustCodeSeq)-10:]
		newCustCode := AppCustCode + com[0].CompanyCode + t.Format("200601") + appCustCodeSeq
		fmt.Println(newCustCode)

		// Convert type for input to customerData
		birthDate, err := time.Parse("2006-01-02 15:04:05", item.CustomerBirthDate)
		if err != nil {
			fmt.Println(err)
		}

		idType, err := strconv.ParseInt(item.CustomerIDType, 10, 8)
		if err != nil {
			fmt.Println(err)
		}

		drawDownDate, err := time.Parse("2006-01-02 15:04:05", item.LoanTglPk)
		if err != nil {
			fmt.Println(err)
		}

		LoanTglPkChanneling, err := time.Parse("2006-01-02 15:04:05", item.LoanTglPkChanneling)
		if err != nil {
			fmt.Println(err)
		}

		// insert to customerData
		customer := models.CustomerDataTab{
			Custcode:          newCustCode,
			PPK:               item.CustomerPpk,
			Name:              item.CustomerName,
			Address1:          item.CustomerAddress1,
			Address2:          item.CustomerAddress2,
			City:              item.CustomerCity,
			Zip:               item.CustomerZip,
			BirthPlace:        item.CustomerBirthPlace,
			BirthDate:         birthDate,
			IdType:            int8(idType),
			IdNumber:          item.CustomerIDNumber,
			MobileNo:          item.CustomerMobileNo,
			DrawdownDate:      drawDownDate,
			TglPkChanneling:   LoanTglPkChanneling,
			MotherMaidenName:  item.CustomerMotherMaidenName,
			ChannelingCompany: item.ScCompany,
			ApprovalStatus:    "9",
		}
		r.db.Create(&customer)

		// Convert type for input to loanData
		LoanOtr, err := strconv.ParseFloat(item.LoanOtr, 64)
		if err != nil {
			fmt.Println(err)
		}

		LoanDownPayment, err := strconv.ParseFloat(item.LoanDownPayment, 64)
		if err != nil {
			fmt.Println(err)
		}

		LoanLoanAmountChanneling, err := strconv.ParseFloat(item.LoanLoanAmountChanneling, 64)
		if err != nil {
			fmt.Println(err)
		}

		LoanInterestFlatChanneling, err := strconv.ParseFloat(item.LoanInterestFlatChanneling, 32)
		if err != nil {
			fmt.Println(err)
		}

		LoanInterestEffectiveChanneling, err := strconv.ParseFloat(item.LoanInterestEffectiveChanneling, 32)
		if err != nil {
			fmt.Println(err)
		}

		LoanEffectivePaymentType, err := strconv.ParseInt(item.LoanEffectivePaymentType, 10, 8)
		if err != nil {
			fmt.Println(err)
		}

		LoanMonthlyPaymentChanneling, err := strconv.ParseFloat(item.LoanMonthlyPaymentChanneling, 64)
		if err != nil {
			fmt.Println(err)
		}

		// insert to loanData
		loan := models.LoanDataTab{
			Custcode:             newCustCode,
			Branch:               item.ScBranchCode,
			OTR:                  LoanOtr,
			DownPayment:          LoanDownPayment,
			LoanAmount:           LoanLoanAmountChanneling,
			LoanPeriod:           item.LoanLoanPeriodChanneling,
			InterestType:         0,
			InterestFlat:         float32(LoanInterestFlatChanneling),
			InterestEffective:    float32(LoanInterestEffectiveChanneling),
			EffectivePaymentType: int8(LoanEffectivePaymentType),
			AdminFee:             0,
			MonthlyPayment:       LoanMonthlyPaymentChanneling,
			InputDate:            t,
			LastModified:         t,
			ModifiedBy:           "",
			InputDate2:           t,
			InputBy:              "",
			LastModified2:        t,
			ModifiedBy2:          "",
		}
		r.db.Create(&loan)
	}

	return sc, nil
}
