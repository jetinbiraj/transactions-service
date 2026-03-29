package accounts

import (
	"errors"
	"reflect"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestNewService(t *testing.T) {
	type args struct {
		repository Repository
	}
	tests := []struct {
		name    string
		args    args
		want    Service
		wantErr bool
	}{
		{
			name:    "repository is nil",
			want:    nil,
			wantErr: true,
		},
		{
			name: "repository is not nil",
			args: args{
				repository: &memoryStore{},
			},
			want: &service{
				repository: &memoryStore{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewService(tt.args.repository)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_CreateAccount(t *testing.T) {
	type fields struct {
		mockRepoFn func(mockRepository *MockRepository)
	}
	type args struct {
		accountRequest Account
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "account creation fail due to save error",
			fields: fields{func(mockRepository *MockRepository) {
				mockRepository.EXPECT().Save(Account{}).Return(
					errors.New("save error"),
				)
			}},
			args: args{
				accountRequest: Account{},
			},
			wantErr: true,
		},
		{
			name: "account creation success",
			fields: fields{func(mockRepository *MockRepository) {
				mockRepository.EXPECT().Save(Account{}).Return(
					nil,
				)
			}},
			args: args{
				accountRequest: Account{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := NewMockRepository(ctrl)

			if tt.fields.mockRepoFn != nil {
				tt.fields.mockRepoFn(mockRepo)
			}

			s := &service{
				repository: mockRepo,
			}
			if err := s.CreateAccount(tt.args.accountRequest); (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetAccount(t *testing.T) {
	const testAccountID = int64(1)
	const testDocumentNumber = "12345"
	type fields struct {
		mockRepoFn func(mockRepository *MockRepository)
	}
	type args struct {
		accountId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    AccountInformationResponse
		wantErr bool
	}{
		{
			name: "error from repo to get by id",
			fields: fields{func(mockRepository *MockRepository) {
				mockRepository.EXPECT().GetById(testAccountID).Return(&Account{}, errors.New("error"))
			}},
			args:    args{accountId: 1},
			wantErr: true,
		},
		{
			name: "get by id success",
			fields: fields{func(mockRepository *MockRepository) {
				mockRepository.EXPECT().GetById(testAccountID).Return(&Account{DocumentNumber: testDocumentNumber}, nil)
			}},
			args:    args{accountId: 1},
			wantErr: false,
			want: AccountInformationResponse{
				AccountId:      1,
				DocumentNumber: testDocumentNumber,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockRepo := NewMockRepository(ctrl)

			if tt.fields.mockRepoFn != nil {
				tt.fields.mockRepoFn(mockRepo)
			}

			s := &service{
				repository: mockRepo,
			}
			got, err := s.GetAccount(tt.args.accountId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}
