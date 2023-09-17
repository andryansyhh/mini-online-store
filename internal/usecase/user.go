package usecase

import (
	"errors"
	"fmt"
	"mini-online-store/helpers"
	"mini-online-store/internal/domain/dto"
	"mini-online-store/internal/domain/models"
	"mini-online-store/internal/repository"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	RegisterUser(req *dto.RegisterUser) error
	Login(req *dto.LoginUser) (*dto.LoginResponse, error)
	GetToken(userModel *models.User) (*dto.UserTokenDto, error)
	LoginCreateToken(user *models.User) (*dto.UserTokenDto, error)
	Auth(req *dto.AuthRequest) (*dto.AuthResponse, error)
}

type userUsecase struct {
	userRepository      repository.UserRepository
	userTokenRepository repository.UserTokenRepository
}

func NewUserUsecase(userRepo repository.UserRepository, userTknRepo repository.UserTokenRepository) UserUsecase {
	return &userUsecase{
		userRepository:      userRepo,
		userTokenRepository: userTknRepo,
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

func (u *userUsecase) Login(req *dto.LoginUser) (*dto.LoginResponse, error) {
	userRes, err := u.userRepository.Login(req.Email)
	if err != nil {
		return nil, errors.New("error login")
	}
	if userRes.Uuid == "" {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userRes.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	tkn, err := u.GetToken(userRes)
	if err != nil {
		return nil, errors.New("error get token")
	}

	return &dto.LoginResponse{
		Token:        tkn.Token,
		RefreshToken: tkn.RefreshToken,
		ExpiredToken: tkn.RefreshTokenExpiredAt,
		UserUuid:     userRes.Uuid,
	}, nil
}

func (u *userUsecase) GetToken(userModel *models.User) (*dto.UserTokenDto, error) {
	existingToken, err := u.userTokenRepository.GetLastToken(userModel.Uuid)
	if err != nil {
		return nil, errors.New("error get last token")
	}
	if existingToken != nil {
		return existingToken, nil
	}
	tkn, err := u.LoginCreateToken(userModel)
	if err != nil {
		return nil, errors.New("error create token")
	}
	return tkn, nil
}

func (u *userUsecase) LoginCreateToken(user *models.User) (*dto.UserTokenDto, error) {
	modelUserToken := models.UserToken{
		UserUuid: user.Uuid,
		Token: helpers.HashToken(fmt.Sprintf("%s%s%s",
			user.Uuid,
			uuid.New().String(),
			time.Now().String())),
		RefreshToken: helpers.HashToken(fmt.Sprintf("%s%s%s",
			user.Uuid,
			uuid.New().String(),
			time.Now().String())),
		TokenExpiredAt:        time.Now().AddDate(0, 0, 30),
		RefreshTokenExpiredAt: time.Now().AddDate(1, 0, 0),
	}
	tkn, err := u.userTokenRepository.CreateUserToken(&modelUserToken)

	if err != nil {
		return nil, errors.New("failed to create user token")
	}
	return tkn, nil
}

func (u *userUsecase) Auth(req *dto.AuthRequest) (*dto.AuthResponse, error) {
	userTkn, err := u.userTokenRepository.FindByToken(req.Token)
	if err != nil {
		return nil, errors.New("unauthorized")
	}
	return &dto.AuthResponse{
		UserUuid: userTkn.UserUuid,
		Email:    userTkn.User.Email,
		Phone:    userTkn.User.Phone,
	}, nil
}
