package usecase

import (
	"fmt"
	"time"
	"user_service/domain"
	"user_service/infrastructure/database/models"
	"user_service/infrastructure/repository"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var TOKEN_KEY = "THIS IS TOKEN KEY"
func generateToken(u *models.User) (string, error) {
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
	Login(user *domain.User) (string, error)
}


func (u *UserUsecase) Login(user *domain.User) (string, error) {
	userModel, err := u.UserRepositoryI.FetchUserByEmail(user.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	tokenString, err := generateToken(userModel)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}