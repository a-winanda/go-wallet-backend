package service

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
	"errors"
	"fmt"
	"time"
)

type TransactionServices interface {
	TopUpWallet(entity.Transaction) error
	TransferWallet(e entity.Transaction) error
}

type transactionServicesImplementation struct {
	transactionRepository repository.TransactionRepository
	userRepo              repository.UserRepository
}

func NewTransactionServices(tr repository.TransactionRepository, ur repository.UserRepository) TransactionServices {
	return &transactionServicesImplementation{
		transactionRepository: tr,
		userRepo:              ur,
	}
}

func (t *transactionServicesImplementation) TopUpWallet(e entity.Transaction) error {

	if e.FundID == 0 {
		return errors.New("fund id cannot be empty")
	}

	if e.Amount < 50000 || e.Amount > 50000000 {
		return errors.New("please input amount between 50.000 and 50 million")
	}

	//fundSource := t.userRepo.
	wallet, err := t.userRepo.GetWalletByUID(e.SourceID)
	if err != nil {
		return err
	}

	e.WalletNumber = wallet.WalletNumber
	e.TransactionType = "Top Up"
	e.Description = fmt.Sprintf("Top Up from %d", e.FundID)
	e.CreatedAt = time.Now()

	err = t.transactionRepository.CreateTransaction(&e)
	if err != nil {
		return err
	}
	t.userRepo.AddWalletBalance(e.WalletNumber, e.Amount)
	return nil
}

func (t *transactionServicesImplementation) TransferWallet(e entity.Transaction) error {

	source, err := t.userRepo.GetWalletByUID(e.SourceID)
	if err != nil {
		return err
	}

	target, err := t.userRepo.GetWalletByUID(e.TargetID)
	if err != nil {
		return err
	}

	if e.Amount < 1000 || e.Amount > 50000000 {
		return errors.New("please input amount between 1000 and 50 million")
	}

	if source.Balance < e.Amount {
		return errors.New("insufficient balance on your wallet")
	}

	if e.Description == "" || len(e.Description) > 35 {
		return errors.New("description need to be less than 35 characters")
	}

	e.WalletNumber = source.WalletNumber
	e.TransactionType = "Transfer"
	e.CreatedAt = time.Now()

	err = t.transactionRepository.CreateTransaction(&e)
	if err != nil {
		return err
	}

	t.userRepo.AddWalletBalance(target.WalletNumber, e.Amount)
	t.userRepo.ReduceWalletBalance(source.WalletNumber, e.Amount)

	return nil
}
