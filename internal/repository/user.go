package repository

import (
	"mini-online-store/internal/domain/dto"
	"mini-online-store/internal/domain/models"

	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

type UserRepository interface {
	CreateUser(userToCreate *dto.RegisterUser) error
	Login(email string) (*models.User, error)
}

func NewUserRepository(db *gorm.DB) User {
	return User{
		db: db,
	}
}

func (u *User) CreateUser(userToCreate *dto.RegisterUser) error {
	if err := u.db.Table("user").Debug().Create(&userToCreate).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Login(email string) (*models.User, error) {
	var res models.User
	if err := u.db.Table("user").Debug().
		Where("email = ? and deleted_at is null", email).
		First(&res).
		Error; err != nil && err.Error() != "record not found" {
		return nil, err
	}
	return &res, nil
}
