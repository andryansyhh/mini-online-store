package usecase

import (
	"errors"
	"mini-online-store/helpers"
	"mini-online-store/internal/domain/dto"
	"mini-online-store/internal/repository"
	"time"

	"github.com/google/uuid"
)

type UserUsecase interface {
	RegisterUser(req *dto.RegisterUser) error
}

type userUsecase struct {
	userRepository repository.UserRepository
	// inquiryRepository             repository.InquiryRepository
	// paymentNotificationRepository repository.PaymentNotificationRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}

func (u *userUsecase) RegisterUser(req *dto.RegisterUser) error {
	if req.Email == "" || req.Password == "" {
		return errors.New("username or password cannot be empty")
	}

	if req.Fullname == "" {
		return errors.New("fulllname cannot be empty")
	}

	userRes, err := u.userRepository.Login(req.Email)
	if err != nil {
		return errors.New("error login")
	}

	// Payload validation for email and phone
	if req.Email == userRes.Email {
		return errors.New("email already registered")
	}

	if req.Phone == userRes.Phone {
		return errors.New("phone already registered")
	}

	// check password
	isMatch := helpers.CheckPassword(req.Password)
	if !isMatch {
		return errors.New("password at least 8 characters Have at least 1 number, 1 symbol Include both Upper case and Lower case characters")
	}

	hash, err := helpers.NewHashPassword(req.Password)
	if err != nil {
		return errors.New("error hasing")
	}

	userToCreate := &dto.RegisterUser{
		Uuid:      uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Email:     req.Email,
		Phone:     req.Phone,
		Fullname:  req.Fullname,
		Password:  hash,
		Address:   req.Address,
	}
	err = u.userRepository.CreateUser(userToCreate)
	if err != nil {
		return errors.New("failed create user")

	}
	return nil
}
