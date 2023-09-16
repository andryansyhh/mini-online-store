package models

type User struct {
	BaseModelUuid
	Fullname string `gorm:"fullname"`
	Email    string `gorm:"email"`
	Phone    string `gorm:"fullname"`
	Password string `gorm:"password"`
	Address  string `gorm:"address"`
}

func (u User) TableName() string {
	return "user"
}