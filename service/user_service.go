package service

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
	"assignment-golang-backend/utils"
	"errors"
	"fmt"
	"time"
)

type UserServices interface {
	GetAllUser() ([]*entity.User, error)
	LoginUser(email, password string) error
	RegisterUser(e entity.User) error
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

func (u *userSevicesImplementation) GetAllUser() ([]*entity.User, error) {

	ul, err := u.repository.GetAllUser()
	if err != nil {
		return nil, errors.New("user database empty")
	}
	return ul, err
}

func (u *userSevicesImplementation) LoginUser(email, password string) error {

	ul, err := u.repository.GetUserByEmail(email)
	fmt.Printf("ul: %v\n", ul)
	if err != nil {
		return err
	}

	//fmt.Println(utils.ComparePassword(ul.Password, password))
	if !utils.ComparePassword(ul.Password, password) {
		return errors.New("wrong password")
	}

	return nil
}
