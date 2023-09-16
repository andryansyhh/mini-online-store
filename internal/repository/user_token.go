package repository

import (
	"mini-online-store/internal/domain/dto"
	"mini-online-store/internal/domain/models"
	"time"

	"gorm.io/gorm"
)

type UserToken struct {
	db *gorm.DB
}

type UserTokenRepository interface {
	CreateUserToken(userToken *models.UserToken) (*dto.UserTokenDto, error)
	GetLastToken(user_uuid string) (*dto.UserTokenDto, error)
}

func NewUserTokenRepository(db *gorm.DB) UserToken {
	return UserToken{
		db: db,
	}
}

func (u *UserToken) CreateUserToken(userToken *models.UserToken) (*dto.UserTokenDto, error) {
	var res dto.UserTokenDto
	err := u.db.Table("user_token").Debug().Create(&userToken).Error
	if err != nil {
		return nil, err
	}

	err = u.db.Table("user_token").
		Where("uuid = ?", userToken.Uuid).First(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *UserToken) GetLastToken(user_uuid string) (*dto.UserTokenDto, error) {
	var token dto.UserTokenDto
	err := u.db.Table("user_token").Debug().
		Where("user_uuid = ? AND token_expired_at > ?", user_uuid, time.Now()).
		Order("created_at DESC").
		First(&token).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &token, nil
}
