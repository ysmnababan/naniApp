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
	FetchUserByPhone(phone_number string) (*domain.User, error)
	FetchUserByEmail(email string) (*domain.User, error)
	FetchUserByID(user_id string) (*domain.User, error)
	IsUserExist(email, phone_number string) error
	CreateUser(user *domain.User) error
	UpdateUser(user *domain.User) error

	FetchAllContact(user_id string) ([]*domain.Phonebook, error)
	CreateContact(contact *domain.Phonebook) error
	UpdateNickname(phonebook_id, nickname string) error
	IsContactUnique(user_id, contact_id string) (bool, error)
}

func (r *Repo) IsContactUnique(user_id, contact_id string) (bool, error) {
	res := r.DB.Where("user_id=? and contact_id=?", user_id, contact_id).First(&models.Phonebook{})
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return true, nil
	}
	if res.Error != nil {
		return false, res.Error
	}
	return false, nil
}

func (r *Repo) UpdateNickname(phonebook_id, nickname string) error {
	phonebook := models.Phonebook{}
	res := r.DB.First(&phonebook, phonebook_id)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return helper.ErrNoData
		}
		return res.Error
	}
	phonebook.Nickname = nickname
	res = r.DB.Save(&phonebook)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *Repo) CreateContact(contact *domain.Phonebook) error {
	phonebook := models.Phonebook{
		PhonebookID: contact.PhonebookID,
		UserID:      contact.UserID,
		ContactID:   contact.ContactID,
		Nickname:    contact.Nickname,
	}

	res := r.DB.Create(&phonebook)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *Repo) FetchAllContact(user_id string) ([]*domain.Phonebook, error) {
	phonebooks := []models.Phonebook{}
	res := r.DB.Where("user_id=?", user_id).Find(&phonebooks)
	if res.Error != nil {
		return nil, res.Error
	}

	// change model to domain
	out := []*domain.Phonebook{}
	for i := range phonebooks {
		contact := &domain.Phonebook{
			PhonebookID: phonebooks[i].PhonebookID,
			UserID:      phonebooks[i].UserID,
			ContactID:   phonebooks[i].ContactID,
			Nickname:    phonebooks[i].Nickname,
		}

		out = append(out, contact)
	}
	return out, nil
}

func (r *Repo) FetchUserByPhone(phone_number string) (*domain.User, error) {
	user := models.User{}
	res := r.DB.Where("phone_number=?", phone_number).First(&user)
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
	res := r.DB.Where("phone_number=? or email=?", phone_number, email).First(&user)
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
