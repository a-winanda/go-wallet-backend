package repository

import (
	"assignment-golang-backend/entity"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateTransaction(e *entity.Transaction) error
	GetAllTransactionDefault(int, entity.TransactionRequest) ([]*entity.Transaction, error)
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

func (t *transactionRepositoryImplementation) GetAllTransactionDefault(uid int, e entity.TransactionRequest) ([]*entity.Transaction, error) {

	var tl []*entity.Transaction

	order := fmt.Sprintf("%s %s", e.SortByEntity, e.SortOrder)

	err := t.db.Preload("Fund").Where("source_id = ? AND description ILIKE ?", uid, e.DescriptionRequest).Order(order).Limit(e.Limit).Find(&tl).Error

	if err != nil {
		return nil, err
	}

	if tl == nil {
		return nil, errors.New("transaction list is empty")
	}

	return tl, nil
}
