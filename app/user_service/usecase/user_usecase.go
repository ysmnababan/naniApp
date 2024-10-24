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
	UpdateUserData(user *domain.User) (*domain.User, error)

	GetContacts(user_id string) ([]*domain.Phonebook, error)
	CreateNewContact(phone_number string, contact *domain.Phonebook) error
	EditNickname(nickname string, contact *domain.Phonebook) error
}

func (u *UserUsecase) GetContacts(user_id string) ([]*domain.Phonebook, error) {
	if user_id == "" {
		return nil, helper.ErrInvalidId
	}

	res, err := u.UserRepositoryI.FetchAllContact(user_id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *UserUsecase) CreateNewContact(phone_number string, contact *domain.Phonebook) error {
	if phone_number == "" || contact.UserID == "" {
		return helper.ErrParam
	}

	// fetch user data using phone number
	user, err := u.UserRepositoryI.FetchUserByPhone(phone_number)
	if err != nil {
		// user not found
		return err
	}

	// ensure no duplicate contact id for one user id
	isUnique, err := u.UserRepositoryI.IsContactUnique(contact.UserID, contact.ContactID)
	if err != nil || !isUnique {
		return err
	}

	// generate uuid
	uuidV7 := uuidv7.New()
	contact.PhonebookID = uuidV7.String()

	// populate contact data
	contact.ContactID = user.UserID
	if contact.Nickname == "" {
		// registered name is default name
		contact.Nickname = user.Username
	}

	err = u.UserRepositoryI.CreateContact(contact)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserUsecase) EditNickname(nickname string, phonebook_id string) error {
	if phonebook_id == "" || nickname == "" {
		return helper.ErrParam
	}

	err := u.UserRepositoryI.UpdateNickname(phonebook_id, nickname)
	if err != nil {
		return nil
	}
	return nil
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

func (u *UserUsecase) UpdateUserData(user *domain.User) (*domain.User, error) {
	// fetch data from databasse first,
	// if req data is not empty, then update
	userDB, err := u.UserRepositoryI.FetchUserByID(user.UserID)
	if err != nil {
		return nil, err
	}

	// username cannot be empty
	if user.Username != "" {
		userDB.Username = user.Username
	}
	userDB.Picture_URL = user.Picture_URL

	// update user data
	err = u.UserRepositoryI.UpdateUser(userDB)
	if err != nil {
		return nil, err
	}
	return userDB, nil
}
