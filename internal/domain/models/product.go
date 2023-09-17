package models

type Product struct {
	BaseModelUuid
	CategoryProductUUid string          `gorm:"column:category_product_uuid" json:"category_product_uuid"`
	CategoryProduct     CategoryProduct `gorm:"foreignKey:CategoryProductUUid"`
	Name                string          `gorm:"column:name"`
	Qty                 int             `gorm:"column:qty"`
	Price               float64         `gorm:"column:price"`
}

func (u Product) TableName() string {
	return "product"
}
