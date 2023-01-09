package skalaAngsuran

import (
	"fmt"
	"log"
	"math"
	"project_credit_sinarmas/backend/models"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	GenerateSkalaAngsuran() ([]models.CustomerDataTab, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) GenerateSkalaAngsuran() ([]models.CustomerDataTab, error) {
	var customer []models.CustomerDataTab
	var loan models.LoanDataTab
	res := r.db.Where("approval_status = ?", "9").Find(&customer)
	if res.Error != nil {
		log.Println("Get Data error : ", res.Error)
		return nil, res.Error
	}
	for _, item := range customer {

		res := r.db.Where("custcode = ?", item.Custcode).First(&loan)
		if res.Error != nil {
			log.Println("Get Data error : ", res.Error)
			return nil, res.Error
		}
		loanPeriod, err := strconv.ParseInt(loan.LoanPeriod, 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		osBalanceConvert := strings.TrimPrefix(loan.LoanAmount, "Rp")
		osBalanceConvert = strings.TrimSuffix(osBalanceConvert, "00")
		osBalanceConvert = strings.ReplaceAll(osBalanceConvert, ".", "")
		osBalanceConvert = strings.ReplaceAll(osBalanceConvert, ",", "")
		osBalanceFloat, err := strconv.ParseFloat(osBalanceConvert, 64)
		if err != nil {
			fmt.Println(err)
		}

		monthlyPaymentConvert := strings.TrimPrefix(loan.MonthlyPayment, "Rp")
		monthlyPaymentConvert = strings.TrimSuffix(monthlyPaymentConvert, "00")
		monthlyPaymentConvert = strings.ReplaceAll(monthlyPaymentConvert, ".", "")
		monthlyPaymentConvert = strings.ReplaceAll(monthlyPaymentConvert, ",", "")
		monthlyPaymentFloat, err := strconv.ParseFloat(monthlyPaymentConvert, 64)
		if err != nil {
			fmt.Println(err)
		}

		osBalance := loan.LoanAmount
		monthlyPayment := loan.MonthlyPayment

		// dataSkalaRental := make([]models.SkalaRentalTab, loanPeriod+1)
		timeNow := time.Now()
		for i := 0; i <= int(loanPeriod); i++ {
			if i == 0 {
				Skala := models.SkalaRentalTab{
					Custcode:   item.Custcode,
					Counter:    int8(i),
					Osbalance:  osBalance,
					EndBalance: osBalance,
					DueDate:    timeNow,
					EffRate:    float64(loan.InterestEffective),
					Rental:     monthlyPayment,
					Principle:  "0",
					Interest:   "0",
					Inputdate:  timeNow,
				}
				r.db.Create(&Skala)
			} else {
				interest := math.Floor(osBalanceFloat * float64(loan.InterestEffective) * 30 / 36000)
				// interestString := strconv.FormatFloat(interest, 'f', -1, 64)
				round := int(math.Round(interest))
				interestString := fmt.Sprintf("%d", round)
				// interestString := fmt.Sprintf("%.0f", interest)
				// interestString = strings.ReplaceAll(interestString, ".", ",")
				// interestString = strings.Join(strings.SplitAfterN(interestString[:len(interestString)-3], "", 3), ".") + interestString[len(interestString)-3:]

				principle := monthlyPaymentFloat - interest
				// principleString := strconv.FormatFloat(principle, 'f', -1, 64)
				round = int(math.Round(principle))
				principleString := fmt.Sprintf("%d", round)
				// principleString := fmt.Sprintf("%.0f", principle)
				// principleString = strings.ReplaceAll(principleString, ".", ",")
				// principleString = strings.Join(strings.SplitAfterN(principleString[:len(principleString)-3], "", 3), ".") + principleString[len(principleString)-3:]

				endBalance := osBalanceFloat - principle
				// endBalanceString := strconv.FormatFloat(endBalance, 'f', -1, 64)
				round = int(math.Round(endBalance))
				endBalanceString := fmt.Sprintf("%d", round)
				// endBalanceString := fmt.Sprintf("%.0f", endBalance)
				// endBalanceString = strings.ReplaceAll(endBalanceString, ".", ",")
				// endBalanceString = strings.Join(strings.SplitAfterN(endBalanceString[:len(endBalanceString)-3], "", 3), ".") + endBalanceString[len(endBalanceString)-3:]

				// osBalanceString := strconv.FormatFloat(osBalanceFloat, 'f', -1, 64)
				round = int(math.Round(osBalanceFloat))
				osBalanceString := fmt.Sprintf("%d", round)
				// osBalanceString := fmt.Sprintf("%.0f", osBalanceFloat)
				// osBalanceString = strings.ReplaceAll(osBalanceString, ".", ",")
				// osBalanceString = strings.Join(strings.SplitAfterN(osBalanceString[:len(osBalanceString)-3], "", 3), ".") + osBalanceString[len(osBalanceString)-3:]

				dueDate := timeNow.AddDate(0, i, 0)
				Skala := models.SkalaRentalTab{
					Custcode:   item.Custcode,
					Counter:    int8(i),
					Osbalance:  osBalanceString,
					EndBalance: endBalanceString,
					DueDate:    dueDate,
					EffRate:    float64(loan.InterestEffective),
					Rental:     monthlyPayment,
					Principle:  principleString,
					Interest:   interestString,
					Inputdate:  timeNow,
				}
				r.db.Create(&Skala)
				osBalanceFloat = endBalance
			}
		}

	}
	return customer, nil
}
