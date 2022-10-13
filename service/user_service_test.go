package service

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/mocks"
	"errors"
	"reflect"
	"testing"
)

var TestDummy1 = &entity.User{
	ID:       8,
	Email:    "tomioka@kny.com",
	Password: "passtomioka",
}

var TestDummy2 = &entity.User{
	ID:       3,
	Email:    "sena@eyeshield.com",
	Password: "passsena",
}

var WalletTest1 = &entity.Wallet{
	WalletNumber: 2700,
	Balance:      700000,
	UserID:       8,
}

const tomiokaPass = "$2a$04$cWZ9cRt7n7L/8tm6T392COt9NOytQdiI8bFkf.VCtKIP8A112k.7e"

func Test_userSevicesImplementation_LoginUser(t *testing.T) {
	type fields struct {
		repository mocks.UserRepository
	}
	type args struct {
		email    string
		password string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr error
	}{

		{
			name: "TestLoginFail",
			fields: fields{
				repository: *mocks.NewUserRepository(t),
			},
			args: args{
				email:    TestDummy1.Email,
				password: TestDummy1.Password,
			},
			want:    "",
			wantErr: errors.New("wrong password"),
		},
		{
			name: "TestLoginSuccess",
			fields: fields{
				repository: *mocks.NewUserRepository(t),
			},
			args: args{
				email:    TestDummy1.Email,
				password: TestDummy1.Password,
			},
			want:    "",
			wantErr: nil,
		}, {
			name: "TestLoginNoUser",
			fields: fields{
				repository: *mocks.NewUserRepository(t),
			},
			args: args{
				email:    "emailngasal",
				password: "passngasal",
			},
			want:    "",
			wantErr: errors.New("user not found"),
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			u := NewUserServices(&tt.fields.repository)
			switch i {
			case 0:
				user := new(entity.User)
				user.Password = "1234"
				tt.fields.repository.On("GetUserByEmail", TestDummy1.Email).Return(user, nil)
			case 1:
				user := new(entity.User)
				user.Email = TestDummy1.Email
				user.Password = tomiokaPass
				tt.fields.repository.On("GetUserByEmail", TestDummy1.Email).Return(user, nil)
			case 2:
				user := new(entity.User)
				user.Email = ""
				user.Password = ""
				tt.fields.repository.On("GetUserByEmail", "emailngasal").Return(nil, errors.New("user not found"))
			}

			_, err := u.LoginUser(tt.args.email, tt.args.password)

			if tt.wantErr != nil {
				errCase := err.Error()
				wantString := tt.wantErr.Error()
				if errCase != wantString {
					t.Errorf("userSevicesImplementation.LoginUser() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			}

		})
	}
}

func Test_userSevicesImplementation_GetUserDetails(t *testing.T) {
	type fields struct {
		repository mocks.UserRepository
	}
	type args struct {
		uid int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Wallet
		want1   *entity.User
		wantErr bool
	}{
		{
			name: "TestGetDetail",
			fields: fields{
				repository: *mocks.NewUserRepository(t),
			},
			args: args{
				uid: 8,
			},
			want:    WalletTest1,
			want1:   TestDummy1,
			wantErr: false,
		},
		{
			name: "TestGetNoUser",
			fields: fields{
				repository: *mocks.NewUserRepository(t),
			},
			args: args{
				uid: 0,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
		{
			name: "TestGetErrWalllet",
			fields: fields{
				repository: *mocks.NewUserRepository(t),
			},
			args: args{
				uid: 0,
			},
			want:    nil,
			want1:   nil,
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			u := NewUserServices(&tt.fields.repository)

			switch i {
			case 0:
				tt.fields.repository.On("GetUserByLogin", tt.args.uid).Return(tt.want1, nil)
				tt.fields.repository.On("GetWalletByUID", tt.args.uid).Return(tt.want, nil)
			case 1:
				tt.fields.repository.On("GetUserByLogin", tt.args.uid).Return(tt.want1, errors.New(""))
				tt.fields.repository.On("GetWalletByUID", tt.args.uid).Return(tt.want, nil)
			case 2:
				tt.fields.repository.On("GetUserByLogin", tt.args.uid).Return(tt.want1, nil)
				tt.fields.repository.On("GetWalletByUID", tt.args.uid).Return(nil, errors.New(""))
			}

			got, got1, _ := u.GetUserDetails(tt.args.uid)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userSevicesImplementation.GetUserDetails() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("userSevicesImplementation.GetUserDetails() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_userSevicesImplementation_RegisterUser(t *testing.T) {
	type fields struct {
		repository mocks.UserRepository
	}
	type args struct {
		e entity.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "TestRegiser",
			fields: fields{
				repository: *mocks.NewUserRepository(t),
			},
			args: args{
				e: *TestDummy1,
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUserServices(&tt.fields.repository)

			switch i {
			case 0:
				tt.fields.repository.On("RegisterUser", &tt.args.e).Return(nil)
				tt.fields.repository.On("GenerateWallet", tt.args.e.ID).Return(nil)

			}

			if err := u.RegisterUser(tt.args.e); (err != nil) != tt.wantErr {
				t.Errorf("userSevicesImplementation.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
