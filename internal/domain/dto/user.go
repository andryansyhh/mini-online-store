package dto

import (
	"time"
)

type RegisterUser struct {
	Uuid      string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Fullname  string    `json:"fullname"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}