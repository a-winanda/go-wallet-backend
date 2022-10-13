package service

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/mocks"
	"reflect"
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

var DummyRequest = &entity.TransactionRequest{
	DescriptionRequest: "trf req",
	SortByEntity:       "amount",
	SortOrder:          "desc",
	Limit:              5,
}

var TransactionDummy = &entity.Transaction{
	ID:              1,
	SourceID:        WalletTest1.UserID,
	TargetID:        WalletTest1.UserID,
	FundID:          FundDummy.ID,
	Amount:          75000,
	WalletNumber:    WalletTest1.WalletNumber,
	TransactionType: "Top Up",
	Description:     "Top Up from ",
}

var TransactionDummy2 = &entity.Transaction{
	ID:              2,
	SourceID:        WalletTest2.UserID,
	TargetID:        WalletTest1.UserID,
	FundID:          FundDummy.ID,
	Amount:          75000,
	WalletNumber:    WalletTest2.WalletNumber,
	TransactionType: "Transfer",
	Description:     "Transfer from ",
}

var SliceDummy = []*entity.Transaction{TransactionDummy, TransactionDummy2}

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
		{
			name: "TestTopUpFundIdEmpty",
			fields: fields{
				userRepo:              *mocks.NewUserRepository(t),
				transactionRepository: *mocks.NewTransactionRepository(t),
			},
			args: args{
				e: entity.Transaction{
					SourceID:        WalletTest1.UserID,
					FundID:          0,
					Amount:          75000,
					WalletNumber:    WalletTest1.WalletNumber,
					TransactionType: "Top Up",
					Description:     "Top Up from ",
				},
			},
			wantErr: true,
		},
		{
			name: "TestTopUpAmountError",
			fields: fields{
				userRepo:              *mocks.NewUserRepository(t),
				transactionRepository: *mocks.NewTransactionRepository(t),
			},
			args: args{
				e: entity.Transaction{
					SourceID:        WalletTest1.UserID,
					FundID:          FundDummy.ID,
					Amount:          5000,
					WalletNumber:    WalletTest1.WalletNumber,
					TransactionType: "Top Up",
					Description:     "Top Up from ",
				},
			},
			wantErr: true,
		},
		{
			name: "TestTopUpWalletError",
			fields: fields{
				userRepo:              *mocks.NewUserRepository(t),
				transactionRepository: *mocks.NewTransactionRepository(t),
			},
			args: args{
				e: entity.Transaction{
					SourceID:        WalletTest1.UserID,
					FundID:          FundDummy.ID,
					Amount:          5000,
					WalletNumber:    WalletTest1.WalletNumber,
					TransactionType: "Top Up",
					Description:     "Top Up from ",
				},
			},
			wantErr: true,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tr := NewTransactionServices(&tt.fields.transactionRepository, &tt.fields.userRepo)

			switch i {
			case 0:
				tt.fields.userRepo.On("GetWalletByUID", tt.args.e.SourceID).Return(WalletTest1, nil)
				tt.fields.transactionRepository.On("CreateTransaction", &tt.args.e).Return(nil)
				tt.fields.userRepo.On("AddWalletBalance", tt.args.e.WalletNumber, tt.args.e.Amount).Return(nil)
				tt.fields.userRepo.On("AddWalletBalance", WalletTest2.WalletNumber, tt.args.e.Amount)
				tt.fields.userRepo.On("ReduceWalletBalance", tt.args.e.SourceID, tt.args.e.Amount)
			case 1:
			case 2:
			case 3:
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
		{
			name: "TestTransferAmountError",
			fields: fields{
				userRepo:              *mocks.NewUserRepository(t),
				transactionRepository: *mocks.NewTransactionRepository(t),
			},
			args: args{
				e: entity.Transaction{
					SourceID:        1,
					TargetID:        1,
					FundID:          0,
					Amount:          0,
					WalletNumber:    0,
					TransactionType: "Transfer",
					Description:     "",
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
				tt.fields.userRepo.On("AddWalletBalance", tt.args.e.WalletNumber, tt.args.e.Amount)
			case 1:
				tt.fields.userRepo.On("GetWalletByUID", tt.args.e.SourceID).Return(WalletTest1, nil)
				tt.fields.userRepo.On("GetWalletByUID", tt.args.e.SourceID).Return(WalletTest1, nil)
			case 2:
			}

			if err := tr.TransferWallet(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("transactionServicesImplementation.TransferWallet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_transactionServicesImplementation_GetAllTransactionByLogin(t *testing.T) {
	type fields struct {
		transactionRepository mocks.TransactionRepository
		userRepo              mocks.UserRepository
	}
	type args struct {
		uid int
		e   entity.TransactionRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entity.Transaction
		wantErr bool
	}{
		{
			name: "TestGetTransaction",
			fields: fields{
				transactionRepository: *mocks.NewTransactionRepository(t),
				userRepo:              *mocks.NewUserRepository(t),
			},
			args: args{
				uid: TransactionDummy.ID,
				e:   *DummyRequest,
			},
			want:    SliceDummy,
			wantErr: false,
		},
		{
			name: "TestGetTransactionSortByEmpty",
			fields: fields{
				transactionRepository: *mocks.NewTransactionRepository(t),
				userRepo:              *mocks.NewUserRepository(t),
			},
			args: args{
				uid: TransactionDummy.ID,
				e: entity.TransactionRequest{
					DescriptionRequest: "DESC",
					SortByEntity:       "",
					SortOrder:          "",
					Limit:              0,
				},
			},
			want:    SliceDummy,
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := NewTransactionServices(&tt.fields.transactionRepository, &tt.fields.userRepo)

			switch i {
			case 0:
				tt.fields.transactionRepository.On("GetAllTransactionDefault", tt.args.uid, tt.args.e).Return(tt.want, nil)
			case 1:
			case 2:

			}

			got, err := tr.GetAllTransactionByLogin(tt.args.uid, tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("transactionServicesImplementation.GetAllTransactionByLogin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transactionServicesImplementation.GetAllTransactionByLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
