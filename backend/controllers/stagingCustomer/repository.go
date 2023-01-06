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

func (r *repository) validationFailed(id int64, reason string, item models.StagingCustomer) {
	var sc []models.StagingCustomer

	// change stagingCustomer sc_flag to 8
	r.db.Model(&sc).Where("id = ?", id).Update("sc_flag", "8")

	//insert to stagingError
	stagingError := models.StagingError{
		SeReff:       item.ScReff,
		SeCreateDate: item.ScCreateDate,
		BranchCode:   item.ScBranchCode,
		Company:      item.ScCompany,
		Ppk:          item.CustomerPpk,
		Name:         item.CustomerName,
		ErrorDesc:    reason,
	}
	r.db.Create(&stagingError)
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
			r.validationFailed(item.ID, "CUSTOMER_PPK tidak boleh duplikasi pada table Customer_data_tab", item)
			continue
		}

		// validasi SC_Company terdaftar
		res2 = r.db.Where("company_short_name = ?", item.ScCompany).First(&com)
		if res2.Error != nil {
			r.validationFailed(item.ID, "SC_COMPANY harus terdaftar di Tabel Mst_Company_Tab", item)
			continue
		}

		// validasi SC_BranchCode terdaftar
		res2 = r.db.Where("code = ?", item.ScBranchCode).First(&b)
		if res2.Error != nil {
			r.validationFailed(item.ID, "C_BRANCH_CODE harus terdaftar di Tabel Branch_Tab", item)
			continue
		}

		// validasi TGL_PK di bulan yang sama
		date, err := time.Parse("2006-01-02", item.LoanTglPk)
		if err != nil {
			fmt.Println(err)
		}
		if (int(t.Month()) != int(date.Month())) || (int(t.Year()) != int(date.Year())) {
			r.validationFailed(item.ID, "TGL_PK / DRAWDOWN_DATE tidak boleh berbeda bulan dengan bulan berjalan saat in", item)
			continue
		}

		// validasi CustomerIDType & CustomerIDNumber
		if item.CustomerIDType == "1" && item.CustomerIDNumber == "" {
			r.validationFailed(item.ID, "Jika 'CUSTOMER_ID_TYPE' diisi = 1, maka 'CUSTOMER_ID_NUMBER' harus diisi dan tidak boleh kosong", item)
			continue
		}

		// validasi Nama Debitur no char spesial
		if strings.ContainsAny(item.CustomerName, "!@#$%^&*()_+-=[]{}|\\:;'\"<>,./?") {
			r.validationFailed(item.ID, "NAMA Debitur tidak boleh mengandung karakter special", item)
			continue
		}

		// validasi Vehicle_BPKB kosong
		if item.VehicleBpkb == "" {
			r.validationFailed(item.ID, "VEHICLE_BPKB tidak boleh kosong", item)
			continue
		}

		// validasi Vehicle_STNK kosong
		if item.VehicleStnk == "" {
			r.validationFailed(item.ID, "VEHICLE_STNK tidak boleh kosong", item)
			continue
		}

		// validasi Vehicle_Engine_NO kosong
		if item.VehicleEngineNo == "" {
			r.validationFailed(item.ID, "VEHICLE_ENGINE_NO tidak boleh kosong", item)
			continue
		}

		// validasi Vehicle_Chasis_NO kosong
		if item.VehicleChasisNo == "" {
			r.validationFailed(item.ID, "VEHICLE_CHASIS_NO tidak boleh kosong", item)
			continue
		}

		// validasi Vehicle_Engine_NO Duplikasi
		res2 = r.db.Where("engine_no = ?", item.VehicleEngineNo).First(&v)
		if res2.Error == nil {
			r.validationFailed(item.ID, "VEHICLE_ENGINE_NO tidak boleh duplikasi pada table 'Vihicle_data_Tab'", item)
			continue
		}

		// validasi Vehicle_Chasis_No Duplikasi
		res2 = r.db.Where("chasis_no = ?", item.VehicleChasisNo).First(&v)
		if res2.Error == nil {
			r.validationFailed(item.ID, "VEHICLE_CHASIS_NO tidak boleh duplikasi pada table 'Vihicle_data_Tab'", item)
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

		// insert to loanData
		loan := models.LoanDataTab{
			Custcode:             newCustCode,
			Branch:               item.ScBranchCode,
			OTR:                  item.LoanOtr,
			DownPayment:          item.LoanDownPayment,
			LoanAmount:           item.LoanLoanAmountChanneling,
			LoanPeriod:           item.LoanLoanPeriodChanneling,
			InterestType:         0,
			InterestFlat:         float32(LoanInterestFlatChanneling),
			InterestEffective:    float32(LoanInterestEffectiveChanneling),
			EffectivePaymentType: int8(LoanEffectivePaymentType),
			AdminFee:             "0",
			MonthlyPayment:       item.LoanMonthlyPaymentChanneling,
			InputDate:            t,
			LastModified:         t,
			ModifiedBy:           "",
			InputDate2:           t,
			InputBy:              "",
			LastModified2:        t,
			ModifiedBy2:          "",
		}
		r.db.Create(&loan)

		// Convert type for input to vehicleData
		VehicleBrand, err := strconv.ParseInt(item.VehicleBrand, 10, 8)
		if err != nil {
			fmt.Println(err)
		}

		VehicleStatus, err := strconv.ParseInt(item.VehicleStatus, 10, 8)
		if err != nil {
			fmt.Println(err)
		}

		VehicleDealerID, err := strconv.ParseInt(item.VehicleDealerID, 10, 8)
		if err != nil {
			fmt.Println(err)
		}

		VehicleTglStnk, err := time.Parse("2006-01-02 15:04:05", item.VehicleTglStnk)
		if err != nil {
			fmt.Println(err)
		}

		VehicleTglBpkb, err := time.Parse("2006-01-02 15:04:05", item.VehicleTglBpkb)
		if err != nil {
			fmt.Println(err)
		}

		CollateralTypeID, err := strconv.ParseInt(item.CollateralTypeID, 10, 8)
		if err != nil {
			fmt.Println(err)
		}

		// insert to vehicleData
		vehicle := models.VehicleDataTab{
			Custcode:       newCustCode,
			Brand:          int(VehicleBrand),
			Type:           item.VehicleType,
			Year:           item.VehicleYear,
			Golongan:       1,
			Jenis:          item.VehicleJenis,
			Status:         int8(VehicleStatus),
			Color:          item.VehicleColor,
			PoliceNo:       item.VehiclePoliceNo,
			EngineNo:       item.VehicleEngineNo,
			ChasisNo:       item.VehicleChasisNo,
			Bpkb:           item.VehicleBpkb,
			RegisterNo:     "1",
			Stnk:           item.VehicleStnk,
			StnkAddress1:   "",
			StnkAddress2:   "",
			StnkCity:       "",
			DealerID:       int(VehicleDealerID),
			Inputdate:      t,
			Inputby:        "system",
			Lastmodified:   t,
			Modifiedby:     "system",
			TglStnk:        VehicleTglStnk,
			TglBpkb:        VehicleTglBpkb,
			TglPolis:       t,
			PolisNo:        item.VehiclePoliceNo,
			CollateralID:   CollateralTypeID,
			Ketagunan:      "",
			AgunanLbu:      "",
			Dealer:         item.VehicleDealer,
			AddressDealer1: item.VehicleAddressDealer1,
			AddressDealer2: item.VehicleAddressDealer2,
			CityDealer:     item.VehicleCityDealer,
		}
		r.db.Create(&vehicle)

		r.db.Model(&sc).Where("id = ?", item.ID).Update("sc_flag", "1")
	}

	return sc, nil
}
