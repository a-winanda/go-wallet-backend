package repository

import "gorm.io/gorm"

type TransactionRepository interface {
}

type transactionRepositoryImplementation struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepositoryImplementation{
		db: db,
	}
}

func (t *transactionRepositoryImplementation) TopUpWallet() error {

	return nil
}
