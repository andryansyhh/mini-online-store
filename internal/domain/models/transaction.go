package models

type Transaction struct {
	BaseModelUuid
	ProductUuid string  `gorm:"column:product_uuid" json:"product_uuid"`
	Product     Product `gorm:"foreignKey:ProductUuid"`
	UserUuid    string  `gorm:"column:user_uuid" json:"user_uuid"`
	User        User    `gorm:"foreignKey:UserUuid"`
	Price       float64 `gorm:"column:price"`
}

func (u Transaction) TableName() string {
	return "transaction"
}
