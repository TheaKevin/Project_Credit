package models

type userTab struct {
	Username string `json:"username" gorm:"not null; type: varchar(100); unique"`
	Email    string `json:"email" gorm:"not null; type: varchar(100); unique"`
	Password string `json:"password" gorm:"type: varchar(100)"`
}

func (m *userTab) TableName() string {
	return "user_Tab"
}
