package service

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/mocks"
	"testing"
)

var FundDummy = &entity.Fund{
	ID:         1,
	SourceName: "Bank",
}

var WalletTest2 = &entity.Wallet{
	WalletNumber: 2700,
	Balance:      70000,
	UserID:       4,
}

func Test_transactionServicesImplementation_TopUpWallet(t *testing.T) {
	type fields struct {
		transactionRepository mocks.TransactionRepository
		userRepo              mocks.UserRepository
	}
	type args struct {
		e entity.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestTopUpSuccess",
			fields: fields{
				userRepo:              *mocks.NewUserRepository(t),
				transactionRepository: *mocks.NewTransactionRepository(t),
			},
			args: args{
				e: entity.Transaction{
					SourceID:        WalletTest1.UserID,
					FundID:          FundDummy.ID,
					Amount:          75000,
					WalletNumber:    WalletTest1.WalletNumber,
					TransactionType: "Top Up",
					Description:     "Top Up from ",
				},
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//ur := NewUserServices(&tt.fields.userRepo)
			tr := NewTransactionServices(&tt.fields.transactionRepository, &tt.fields.userRepo)

			switch i {
			case 0:
				tt.fields.userRepo.On("GetWalletByUID", tt.args.e.SourceID).Return(WalletTest1, nil)
				tt.fields.transactionRepository.On("CreateTransaction", &tt.args.e).Return(nil)
				tt.fields.userRepo.On("AddWalletBalance", tt.args.e.WalletNumber, tt.args.e.Amount).Return(nil)
				tt.fields.userRepo.On("AddWalletBalance", WalletTest2.WalletNumber, tt.args.e.Amount)
				tt.fields.userRepo.On("ReduceWalletBalance", tt.args.e.SourceID, tt.args.e.Amount)
			}

			if tr.TopUpWallet(tt.args.e) != nil {
				errBoole := true
				if errBoole != tt.wantErr {
					t.Errorf("transactionSersvicesImplementation.TopUpWallet() error = %v, wantErr %v", errBoole, tt.wantErr)
				}
			}

		})
	}
}

func Test_transactionServicesImplementation_TransferWallet(t *testing.T) {
	type fields struct {
		transactionRepository mocks.TransactionRepository
		userRepo              mocks.UserRepository
	}
	type args struct {
		e entity.Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestTransfer",
			fields: fields{
				userRepo:              *mocks.NewUserRepository(t),
				transactionRepository: *mocks.NewTransactionRepository(t),
			},
			args: args{
				e: entity.Transaction{
					SourceID:        WalletTest1.UserID,
					TargetID:        WalletTest2.UserID,
					FundID:          0,
					Amount:          75000,
					WalletNumber:    WalletTest1.WalletNumber,
					TransactionType: "Transfer",
					Description:     "Transfer from dummy 1 to 2",
				},
			},
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTransactionServices(&tt.fields.transactionRepository, &tt.fields.userRepo)

			switch i {
			case 0:
				tt.fields.userRepo.On("GetWalletByUID", tt.args.e.SourceID).Return(WalletTest1, nil)
				tt.fields.userRepo.On("GetWalletByUID", tt.args.e.TargetID).Return(WalletTest2, nil)
				tt.fields.transactionRepository.On("CreateTransaction", &tt.args.e).Return(nil)
				tt.fields.userRepo.On("AddWalletBalance", tt.args.e.WalletNumber, tt.args.e.Amount).Return(nil)
			}

			if err := tr.TransferWallet(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("transactionServicesImplementation.TransferWallet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
