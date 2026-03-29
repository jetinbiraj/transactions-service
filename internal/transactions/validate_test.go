package transactions

import "testing"

func Test_validateTransaction(t *testing.T) {
	type args struct {
		req CreateTransactionRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "account id is invalid",
			args:    args{req: CreateTransactionRequest{AccountId: -1}},
			wantErr: true,
		},
		{
			name: "operation type is invalid",
			args: args{
				req: CreateTransactionRequest{
					AccountId:       1,
					OperationTypeId: 5,
				}},
			wantErr: true,
		},
		{
			name: "transaction amount is invalid",
			args: args{
				req: CreateTransactionRequest{
					AccountId:       1,
					OperationTypeId: 4,
					Amount:          0.0,
				}},
			wantErr: true,
		},
		{
			name: "positive amount for negative transaction",
			args: args{
				req: CreateTransactionRequest{
					AccountId:       1,
					OperationTypeId: 1,
					Amount:          50.0,
				}},
			wantErr: true,
		},
		{
			name: "negative amount for positive transaction",
			args: args{
				req: CreateTransactionRequest{
					AccountId:       1,
					OperationTypeId: 4,
					Amount:          -50.0,
				}},
			wantErr: true,
		},
		{
			name: "validation success",
			args: args{
				req: CreateTransactionRequest{
					AccountId:       1,
					OperationTypeId: 4,
					Amount:          50.0,
				}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateTransaction(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("validateTransaction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
