package accounts

import (
	"reflect"
	"sync"
	"testing"
)

func TestNewRepository(t *testing.T) {
	tests := []struct {
		name string
		want Repository
	}{
		{
			name: "init memory store with test data",
			want: &memoryStore{
				mu:               sync.RWMutex{},
				currentAccountId: 1,
				accounts: map[int64]*Account{
					1: {
						DocumentNumber: "12345678900",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMemoryStore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMemoryStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryStore_Save(t *testing.T) {
	testDocumentNumber := "12345678900"

	type fields struct {
		currentAccountId int64
		accounts         map[int64]*Account
	}
	type args struct {
		account Account
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantErr            bool
		wantAccountId      int64
		wantAccountsLength int
	}{
		{
			name: "save account",
			fields: fields{
				currentAccountId: 0,
				accounts:         map[int64]*Account{},
			},
			args: args{
				account: Account{DocumentNumber: testDocumentNumber},
			},
			wantErr:            false,
			wantAccountsLength: 1,
			wantAccountId:      1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &memoryStore{
				currentAccountId: tt.fields.currentAccountId,
				accounts:         tt.fields.accounts,
			}

			got, err := r.Save(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.wantAccountId {
				t.Errorf("Save() got = %v, wantAccountId %v", got, tt.wantAccountId)
			}

			if tt.wantAccountsLength != len(r.accounts) {
				t.Errorf("Save() accounts length = %v, wantAccountsLength %v", len(r.accounts), tt.wantAccountsLength)
			}
		})
	}
}

func Test_memoryStore_GetById(t *testing.T) {
	testDocumentNumber := "12345678900"
	type fields struct {
		currentAccountId int64
		accounts         map[int64]*Account
	}
	type args struct {
		accountId int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Account
		wantErr bool
	}{
		{
			name: "account does not exists for the id",
			fields: fields{
				accounts: map[int64]*Account{},
			},
			args: args{
				accountId: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "account exists for the id",
			fields: fields{
				accounts: map[int64]*Account{
					1: {
						DocumentNumber: testDocumentNumber,
					},
				},
			},
			args: args{
				accountId: 1,
			},
			want:    &Account{DocumentNumber: testDocumentNumber},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &memoryStore{
				currentAccountId: tt.fields.currentAccountId,
				accounts:         tt.fields.accounts,
			}
			got, err := r.GetById(tt.args.accountId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
