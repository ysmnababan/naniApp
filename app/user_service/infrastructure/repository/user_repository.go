package repository

import (
	"user_service/domain"
	"user_service/helper"
	"user_service/infrastructure/database/models"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

type UserRepositoryI interface {
	FetchUserByEmail(email string) (*domain.User, error)
	FetchUserByID(user_id string) (*domain.User, error)
	IsUserExist(email, phone_number string) error
	CreateUser(user *domain.User) error
	UpdateUser(user *domain.User) error
}

func (r *Repo) FetchUserByEmail(email string) (*domain.User, error) {
	user := models.User{}
	res := r.DB.Where("email=?", email).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	userDomain := domain.User{
		UserID:      user.UserID,
		Username:    user.Username,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Picture_URL: user.PictureURL,
	}
	return &userDomain, nil
}

func (r *Repo) FetchUserByID(user_id string) (*domain.User, error) {
	user := models.User{}
	res := r.DB.Where("user_id", user_id).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	userDomain := domain.User{
		UserID:      user.UserID,
		Username:    user.Username,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Picture_URL: user.PictureURL,
	}
	return &userDomain, nil
}

func (r *Repo) IsUserExist(email, phone_number string) error {
	user := models.User{}
	res := r.DB.Where("phone_number=? and email=?", phone_number, email).First(&user)
	if res.Error != nil && res.Error != gorm.ErrRecordNotFound {
		return res.Error
	} else if res.Error == nil {
		return helper.ErrUserExists
	}
	return nil
}

func (r *Repo) CreateUser(user *domain.User) error {
	userModel := models.User{
		UserID:      user.UserID,
		Email:       user.Email,
		Username:    user.Username,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		PictureURL:  user.Picture_URL,
	}
	res := r.DB.Create(&userModel)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *Repo) UpdateUser(user *domain.User) error {
	userModel := models.User{
		UserID:      user.UserID,
		Email:       user.Email,
		Username:    user.Username,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		PictureURL:  user.Picture_URL,
	}
	res := r.DB.Save(&userModel)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
