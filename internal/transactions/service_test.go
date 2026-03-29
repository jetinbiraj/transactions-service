package transactions

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

func Test_service_CreateTransaction(t *testing.T) {
	type fields struct {
		mockRepoFn func(mockRepository *MockRepository)
	}
	type args struct {
		transaction Transaction
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "transaction creation fail due to save error",
			fields: fields{func(mockRepository *MockRepository) {
				mockRepository.EXPECT().Save(gomock.Any()).Return(
					errors.New("save error"),
				)
			}},
			args: args{
				transaction: Transaction{},
			},
			wantErr: true,
		},
		{
			name: "transaction creation success",
			fields: fields{func(mockRepository *MockRepository) {
				mockRepository.EXPECT().Save(gomock.Any()).Return(
					nil,
				)
			}},
			args: args{
				transaction: Transaction{},
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
			if err := s.CreateTransaction(tt.args.transaction); (err != nil) != tt.wantErr {
				t.Errorf("CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
