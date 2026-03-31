package transactions

import (
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewRepository(t *testing.T) {
	testTime, err := time.Parse(time.RFC3339Nano, "2020-01-01T10:32:07.7199222Z")
	require.NoError(t, err)

	tests := []struct {
		name string
		want Repository
	}{
		{
			name: "init memory store with test data",
			want: &memoryStore{
				mu:                   sync.RWMutex{},
				currentTransactionId: 4,
				transactions: map[int64]*Transaction{
					1: {
						AccountId:       1,
						OperationTypeId: InstallmentPurchase,
						Amount:          -50.0,
						EventDate:       testTime,
					},
					2: {
						AccountId:       1,
						OperationTypeId: InstallmentPurchase,
						Amount:          -23.5,
						EventDate:       testTime,
					},
					3: {
						AccountId:       1,
						OperationTypeId: InstallmentPurchase,
						Amount:          -18.7,
						EventDate:       testTime,
					},
					4: {
						AccountId:       1,
						OperationTypeId: CreditVoucher,
						Amount:          60.0,
						EventDate:       testTime,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_memoryStore_Save(t *testing.T) {
	type fields struct {
		currentTransactionId int64
		transactions         map[int64]*Transaction
	}
	type args struct {
		transaction Transaction
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantErr            bool
		wantTransactionId  int64
		wantAccountsLength int
	}{
		{
			name: "save transaction",
			fields: fields{
				currentTransactionId: 0,
				transactions:         map[int64]*Transaction{},
			},
			args: args{
				transaction: Transaction{
					AccountId:       1,
					OperationTypeId: Purchase,
					Amount:          -50.0,
					EventDate:       time.Now().UTC(),
				},
			},
			wantErr:            false,
			wantAccountsLength: 1,
			wantTransactionId:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &memoryStore{
				currentTransactionId: tt.fields.currentTransactionId,
				transactions:         tt.fields.transactions,
			}
			got, err := r.Save(tt.args.transaction)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantTransactionId != got {
				t.Errorf("Save() got = %v, wantTransactionId %v", got, tt.wantTransactionId)
			}

			if tt.wantAccountsLength != len(r.transactions) {
				t.Errorf("Save() accounts length = %v, wantAccountsLength %v", len(r.transactions), tt.wantAccountsLength)
			}
		})
	}
}
