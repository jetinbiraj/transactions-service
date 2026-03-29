package accounts

import "testing"

func Test_validateDocumentNumber(t *testing.T) {
	type args struct {
		doc string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty document number",
			args: args{
				doc: "",
			},
			wantErr: true,
		},
		{
			name: "document number not equal to 11",
			args: args{
				doc: "123456789",
			},
			wantErr: true,
		},
		{
			name: "document number contains alphanumeric characters",
			args: args{
				doc: "123456xyz78",
			},
			wantErr: true,
		},
		{
			name: "document number is valid",
			args: args{
				doc: "12345678901",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateDocumentNumber(tt.args.doc); (err != nil) != tt.wantErr {
				t.Errorf("validateDocumentNumber() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateAccountId(t *testing.T) {
	type args struct {
		accountId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty account id",
			args: args{
				accountId: "",
			},
			wantErr: true,
		},
		{
			name: "account id is not parsable to integer",
			args: args{
				accountId: "123abc",
			},
			wantErr: true,
		},
		{
			name: "account id is negative",
			args: args{
				accountId: "-123",
			},
			wantErr: true,
		},
		{
			name: "account id valid",
			args: args{
				accountId: "123",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateAccountId(tt.args.accountId); (err != nil) != tt.wantErr {
				t.Errorf("validateAccountId() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
