package dto

import "time"

type UserTokenDto struct {
	UserUuid              string    `gorm:"column:user_uuid" json:"user_uuid"`
	Token                 string    `gorm:"column:token" json:"token"`
	RefreshToken          string    `gorm:"column:refresh_token;" json:"refresh_token"`
	TokenExpiredAt        time.Time `gorm:"column:token_expired_at;" json:"token_expired_at"`
	RefreshTokenExpiredAt time.Time `gorm:"column:refresh_token_expired_at;" json:"refresh_token_expired_at"`
}

type AuthRequest struct {
	Token string `json:"token"`
}

type AuthResponse struct {
	UserUuid string `json:"user_uuid"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}