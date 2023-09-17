package dto

type ProductResponse struct {
	Uuid                string  `json:"uuid"`
	Name                string  `json:"name"`
	CategoryProductName string  `gorm:"column:name" json:"category_product_name"`
	Qty                 int     `json:"qty"`
	Price               float64 `json:"price"`
}