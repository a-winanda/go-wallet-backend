package repository

import (
	"assignment-golang-backend/entity"
	"errors"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(e *entity.Transaction) error
	GetAllTransaction() ([]*entity.Transaction, error)
}

type transactionRepositoryImplementation struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepositoryImplementation{
		db: db,
	}
}

func (t *transactionRepositoryImplementation) CreateTransaction(e *entity.Transaction) error {

	res := t.db.Create(&e)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (t *transactionRepositoryImplementation) GetAllTransaction() ([]*entity.Transaction, error) {

	var tl []*entity.Transaction

	t.db.Find(&tl)

	if tl == nil {
		return nil, errors.New("transaction list is empty")

	}

	return tl, nil
}
