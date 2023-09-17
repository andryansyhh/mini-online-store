package models

type ShoppingCart struct {
	BaseModelUuid
	ProductUuid string  `gorm:"column:product_uuid" json:"product_uuid"`
	Product     Product `gorm:"foreignKey:ProductUuid"`
	UserUuid    string  `gorm:"column:user_uuid" json:"user_uuid"`
	User        User    `gorm:"foreignKey:UserUuid"`
	Qty         int     `gorm:"column:qty"`
}

func (u ShoppingCart) TableName() string {
	return "shopping_cart"
}