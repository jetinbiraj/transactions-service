package accounts

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"transactions-service/domain"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestNewHandler(t *testing.T) {
	type args struct {
		logEnabled bool
		service    Service
	}
	tests := []struct {
		name    string
		args    args
		want    *Handler
		wantErr bool
	}{
		{
			name:    "service is nil",
			want:    nil,
			wantErr: true,
		},
		{
			name: "service is not nil",
			args: args{
				logEnabled: true,
				service:    &service{},
			},
			want: &Handler{
				logEnabled: true,
				service:    &service{},
			},
			wantErr: false,
		},
		{
			name: "log enabled flag is false",
			args: args{
				logEnabled: false,
				service:    &service{},
			},
			want: &Handler{
				logEnabled: false,
				service:    &service{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewHandler(tt.args.logEnabled, tt.args.service)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHandler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHandler() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandler_CreateAccount(t *testing.T) {
	type fields struct {
		mockServiceFn       func(mockService *MockService)
		requestBodyFilePath string
	}
	tests := []struct {
		name           string
		fields         fields
		wantStatusCode int
	}{
		{
			name:           "request body is nil",
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "request body is invalid",
			fields: fields{
				requestBodyFilePath: "create_transaction_request_invalid.json",
			},
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "error from service call",
			fields: fields{
				requestBodyFilePath: "create_transaction_request_valid.json",
				mockServiceFn: func(mockService *MockService) {
					mockService.EXPECT().CreateAccount(Account{
						DocumentNumber: "12345678901",
					}).Return(errors.New("create account error"))
				},
			},
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name: "account creation success",
			fields: fields{
				requestBodyFilePath: "create_transaction_request_valid.json",
				mockServiceFn: func(mockService *MockService) {
					mockService.EXPECT().CreateAccount(Account{
						DocumentNumber: "12345678901",
					}).Return(nil)
				},
			},
			wantStatusCode: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := NewMockService(ctrl)

			if tt.fields.mockServiceFn != nil {
				tt.fields.mockServiceFn(mockService)
			}

			var body *bytes.Reader
			if tt.fields.requestBodyFilePath != "" {
				fileBytes, err := readJSONFromFile(tt.fields.requestBodyFilePath)
				require.NoError(t, err)
				body = bytes.NewReader(fileBytes)
			} else {
				body = bytes.NewReader([]byte{})
			}

			mux := http.NewServeMux()
			RegisterRoutes(mux, &Handler{service: mockService})
			server := httptest.NewServer(mux)
			defer server.Close()

			resp, err := http.Post(server.URL+"/accounts", "application/json", body)
			require.NoError(t, err)

			if resp.StatusCode != tt.wantStatusCode {
				t.Fatalf("CreateAccount() Status ErrorCode = %v, want %v", resp.StatusCode, tt.wantStatusCode)
			}
		})
	}
}

func TestHandler_GetAccount(t *testing.T) {
	type fields struct {
		mockServiceFn func(mockService *MockService)
	}
	type args struct {
		accountId int
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
		wantResponse   AccountInformationResponse
	}{
		{
			name:           "account id is invalid",
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name: "error from service call",
			args: args{
				accountId: 1,
			},
			fields: fields{func(mockService *MockService) {
				mockService.EXPECT().GetAccount(int64(1)).Return(AccountInformationResponse{}, domain.ErrNotFound)
			}},
			wantStatusCode: http.StatusNotFound,
		},
		{
			name: "get account success",
			args: args{
				accountId: 1,
			},
			fields: fields{func(mockService *MockService) {
				mockService.EXPECT().GetAccount(int64(1)).Return(AccountInformationResponse{
					AccountId:      1,
					DocumentNumber: "12345678901",
				}, nil)
			}},
			wantStatusCode: http.StatusOK,
			wantResponse: AccountInformationResponse{
				AccountId:      1,
				DocumentNumber: "12345678901",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := NewMockService(ctrl)

			if tt.fields.mockServiceFn != nil {
				tt.fields.mockServiceFn(mockService)
			}

			mux := http.NewServeMux()

			RegisterRoutes(mux, &Handler{service: mockService})
			server := httptest.NewServer(mux)
			defer server.Close()

			resp, err := http.Get(server.URL + fmt.Sprintf("/accounts/%v", tt.args.accountId))
			require.NoError(t, err)

			if resp.StatusCode != tt.wantStatusCode {
				t.Fatalf("Status ErrorCode = %v, want %v", resp.StatusCode, tt.wantStatusCode)
				return
			}

			var respBody AccountInformationResponse
			require.NoError(t, json.NewDecoder(resp.Body).Decode(&respBody))

			if !reflect.DeepEqual(respBody, tt.wantResponse) {
				t.Errorf("GetAccount() got = %v, want %v", respBody, tt.wantResponse)
			}
		})
	}
}

func readJSONFromFile(fileName string) ([]byte, error) {
	return os.ReadFile(fmt.Sprintf("./testresources/%s", fileName))
}
