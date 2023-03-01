package models

type UserTab struct {
	Username string `json:"username" gorm:"not null; type: varchar(100); unique"`
	Email    string `json:"email" gorm:"not null; type: varchar(100); unique"`
	Password []byte `json:"password" gorm:"type: varchar(255)"`
}

func (m *UserTab) TableName() string {
	return "user_Tab"
}
