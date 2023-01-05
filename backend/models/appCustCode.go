package models

type AppCustCodeTab struct {
	ID    int64 `json:"id" gorm:"not null; type: bigint"`
	Value int   `json:"value" gorm:"type: bigint"`
}

func (m *AppCustCodeTab) TableName() string {
	return "appCustCode_Tab"
}
