package dto

import "time"

type AddToCart struct {
	Uuid        string    `json:"uuid"`
	CreatedAt   time.Time `json:"created_at"`
	UserUuid    string    `json:"user_uuid"`
	ProductUuid string    `json:"product_uuid"`
	Qty         int       `json:"qty"`
}

type CartResponse struct {
	Uuid        string  `json:"uuid"`
	UserUuid    string  `json:"user_uuid"`
	ProductName string  `gorm:"column:name" json:"product_name"`
	Price       float64 `json:"price"`
	Qty         int     `json:"qty"`
}