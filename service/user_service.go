package service

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
	"assignment-golang-backend/utils"
	"errors"
	"time"
)

type UserServices interface {
	LoginUser(email, password string) (string, error)
	RegisterUser(e entity.User) error
	GetUserDetails(int) (*entity.Wallet, *entity.User, error)
}

type userSevicesImplementation struct {
	repository repository.UserRepository
}

func NewUserServices(repository repository.UserRepository) UserServices {
	return &userSevicesImplementation{
		repository: repository,
	}
}

func (u *userSevicesImplementation) RegisterUser(e entity.User) error {

	if e.Email == "" {
		return errors.New("email cant be empty")
	}

	if e.Password == "" {
		return errors.New("password cant be empty")
	}

	createTime := time.Now()
	hashPassword, err := utils.HashAndSalt(e.Password)
	if err != nil {
		return err
	}

	e.Password = hashPassword
	e.CreatedAt = createTime
	e.UpdatedAt = time.Time{}

	err = u.repository.RegisterUser(&e)
	if err != nil {
		return err
	}

	u.repository.GenerateWallet(e.ID)

	return nil
}

func (u *userSevicesImplementation) GetUserDetails(uid int) (*entity.Wallet, *entity.User, error) {
	ud, err := u.repository.GetUserByLogin(uid)
	if err != nil {
		return nil, nil, err
	}

	w, err := u.repository.GetWalletByUID(uid)
	if err != nil {
		return nil, nil, err
	}

	return w, ud, nil
}

func (u *userSevicesImplementation) LoginUser(email, password string) (string, error) {

	ul, err := u.repository.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if !utils.ComparePassword(ul.Password, password) {
		return "", errors.New("wrong password")
	}
	token, _ := utils.GenerateToken(uint(ul.ID))

	return token, nil
}
