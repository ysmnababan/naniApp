package repository

import (
	"user_service/infrastructure/database/models"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

type UserRepositoryI interface {
	FetchUserByEmail(email string) (*models.User, error)
}

func (r *Repo) FetchUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	res := r.DB.First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
