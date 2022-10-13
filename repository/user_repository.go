package repository

import (
	"assignment-golang-backend/entity"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(string) (*entity.User, error)
	RegisterUser(e *entity.User) error
	GenerateWallet(int) error
	GetWalletByUID(int) (*entity.Wallet, error)
	AddWalletBalance(int, int) error
	ReduceWalletBalance(int, int) error
	GetUserByLogin(int) (*entity.User, error)
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

func (u *userRepositoryImplementation) GetUserByLogin(uid int) (*entity.User, error) {

	var user *entity.User

	err := u.db.Where("id = ?", uid).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepositoryImplementation) GenerateWallet(id int) error {

	NewWallet := &entity.Wallet{
		UserID:  id,
		Balance: 0,
	}

	err := u.db.Omit("WalletNumber").Create(&NewWallet).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepositoryImplementation) GetWalletByUID(uid int) (*entity.Wallet, error) {

	var w *entity.Wallet

	err := u.db.Where("user_id = ?", uid).First(&w).Error
	if err != nil {
		return nil, err
	}

	return w, nil
}

func (u *userRepositoryImplementation) AddWalletBalance(wn, amount int) error {

	var wallet *entity.Wallet

	err := u.db.Where("wallet_number = ?", wn).First(&wallet).Error

	if err != nil {
		return err
	}

	u.db.Model(&wallet).Where("wallet_number = ?", wn).UpdateColumn("balance", gorm.Expr("balance + ?", amount))
	fmt.Printf("wallet.Balance: %v\n", wallet)

	return nil
}

func (u *userRepositoryImplementation) ReduceWalletBalance(wn, amount int) error {

	var wallet *entity.Wallet

	err := u.db.Where("wallet_number = ?", wn).First(&wallet).Error

	if err != nil {
		return err
	}

	u.db.Model(&wallet).Where("wallet_number = ?", wn).UpdateColumn("balance", gorm.Expr("balance - ?", amount))

	return nil
}
