package repository

import (
	"assignment-golang-backend/entity"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser() ([]*entity.User, error)
	GetUserByEmail(email string) (*entity.User, error)
	RegisterUser(e *entity.User) error
}

type userRepositoryImplementation struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImplementation{
		db: db,
	}
}

func (u *userRepositoryImplementation) RegisterUser(e *entity.User) error {

	result := u.db.Create(&e)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *userRepositoryImplementation) GetUserByEmail(email string) (*entity.User, error) {

	var user *entity.User

	res := u.db.Where("email = ?", email).Find(&user).Error
	if res != nil {
		return nil, res
	}
	return user, nil
}

func (u *userRepositoryImplementation) GetAllUser() ([]*entity.User, error) {

	var users []*entity.User

	u.db.Find(&users)

	if users == nil {
		return nil, errors.New("user database is empty")
	}
	return users, nil
}
