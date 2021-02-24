package model

type User struct {
	ID           uint   `gorm:"primary_key;auto_increment"`
	Username     string `gorm:"size:64,uniqueIndex"`
	PasswordHash string `gorm:"size:128"`
}

func (User) TableName() string {
	return "users"
}
