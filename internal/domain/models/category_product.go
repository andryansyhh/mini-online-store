package models

type CategoryProduct struct {
	BaseModelUuid
	Name string `gorm:"column:name"`
}

func (u CategoryProduct) TableName() string {
	return "category_product"
}