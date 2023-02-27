package models

type UserTab struct {
	Username string `json:"username" gorm:"not null; type: varchar(100); unique"`
	Email    string `json:"email" gorm:"not null; type: varchar(100); unique"`
	Password string `json:"-" gorm:"type: varchar(100)"`
}

func (m *UserTab) TableName() string {
	return "user_Tab"
}
