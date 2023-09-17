package dto

import "time"

type TrxResponse struct {
	Uuid        string    `json:"uuid"`
	CreatedAt   time.Time `json:"created_at"`
	ProductUuid string    `json:"product_uuid"`
	UserUuid    string    `json:"user_uuid"`
	Price       float64   `gorm:"column:price"`
}
