package usecase

import (
	"fmt"
	"time"
	"user_service/domain"
	"user_service/helper"
	"user_service/infrastructure/repository"

	"github.com/golang-jwt/jwt"
	"github.com/samborkent/uuidv7"
	"golang.org/x/crypto/bcrypt"
)

var TOKEN_KEY = "THIS IS TOKEN KEY"

func generateToken(u *domain.User) (string, error) {
	// create the payload
	payload := jwt.MapClaims{
		"id":    u.UserID,
		"email": u.Username,
		"exp":   time.Now().Add(time.Hour * 48).Unix(),
	}

	// define the method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// get token string
	tokenString, err := token.SignedString([]byte(TOKEN_KEY))
	if err != nil {
		return "", fmt.Errorf("error when creating token: %v", err)
	}

	return tokenString, nil
}

type UserUsecase struct {
	repository.UserRepositoryI
}

type UserUsecaseI interface {
	Login(email, password string) (string, error)
	Register(user *domain.User) error
	GetUserData(user_id string) (*domain.User, error)
	UpdateUserData(user *domain.User) error
}

func (u *UserUsecase) Login(email, password string) (string, error) {
	userDomain, err := u.UserRepositoryI.FetchUserByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDomain.Password), []byte(password))
	if err != nil {
		return "", err
	}

	tokenString, err := generateToken(userDomain)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func (u *UserUsecase) Register(user *domain.User) error {
	err := u.UserRepositoryI.IsUserExist(user.Email, user.PhoneNumber)
	if err != nil {
		return err
	}

	// generate uuid
	uuidV7 := uuidv7.New()
	user.UserID = uuidV7.String()

	// hash password
	hashedpwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return helper.ErrParam
	}
	user.Password = string(hashedpwd)

	err = u.UserRepositoryI.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) GetUserData(user_id string) (*domain.User, error) {
	user, err := u.UserRepositoryI.FetchUserByID(user_id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecase) UpdateUserData(user *domain.User) error {
	// fetch data from databasse first,
	// if req data is not empty, then update
	userDomain, err := u.UserRepositoryI.FetchUserByID(user.UserID)
	if err != nil {
		return err
	}

	// username cannot be empty
	if user.Username != "" {
		userDomain.Username = user.Username
	}
	userDomain.Picture_URL = user.Picture_URL

	// update user data
	err = u.UserRepositoryI.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}
