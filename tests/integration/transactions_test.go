package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"transactions-service/internal/transactions"

	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// POST /transactions — success cases (201)
// ---------------------------------------------------------------------------

func TestTransactionCreation_SuccessCases(t *testing.T) {
	tests := []struct {
		name string
		req  transactions.CreateTransactionRequest
	}{
		{
			name: "purchase (op=1, negative amount)",
			req:  transactions.CreateTransactionRequest{AccountId: 1, OperationTypeId: 1, Amount: -50.0},
		},
		{
			name: "installment purchase (op=2, negative amount)",
			req:  transactions.CreateTransactionRequest{AccountId: 1, OperationTypeId: 2, Amount: -23.5},
		},
		{
			name: "withdrawal (op=3, negative amount)",
			req:  transactions.CreateTransactionRequest{AccountId: 1, OperationTypeId: 3, Amount: -18.7},
		},
		{
			name: "credit voucher (op=4, positive amount)",
			req:  transactions.CreateTransactionRequest{AccountId: 1, OperationTypeId: 4, Amount: 60.0},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			body, err := json.Marshal(tc.req)
			require.NoError(t, err)

			resp, err := http.Post(testServer.URL+"/transactions", "application/json", bytes.NewReader(body))
			require.NoError(t, err)
			defer resp.Body.Close()

			require.Equal(t, http.StatusCreated, resp.StatusCode)

			var txResp transactions.TransactionResponse
			require.NoError(t, json.NewDecoder(resp.Body).Decode(&txResp))
			require.Positive(t, txResp.TransactionId)
			require.Equal(t, tc.req.AccountId, txResp.AccountId)
			require.Equal(t, tc.req.OperationTypeId, txResp.OperationTypeId)
			require.Equal(t, tc.req.Amount, txResp.Amount)
		})
	}
}

// ---------------------------------------------------------------------------
// POST /transactions — invalid cases (400)
// ---------------------------------------------------------------------------

func TestTransactionCreation_InvalidCases(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantStatus int
	}{
		{
			name:       "malformed JSON",
			body:       `{invalid-json`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "zero account_id",
			body:       `{"account_id":0,"operation_type_id":1,"amount":-50.0}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "negative account_id",
			body:       `{"account_id":-1,"operation_type_id":1,"amount":-50.0}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "operation_type_id zero (out of range)",
			body:       `{"account_id":1,"operation_type_id":0,"amount":-50.0}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "operation_type_id 5 (out of range)",
			body:       `{"account_id":1,"operation_type_id":5,"amount":-50.0}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "zero amount",
			body:       `{"account_id":1,"operation_type_id":1,"amount":0}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "purchase (op=1) with positive amount",
			body:       `{"account_id":1,"operation_type_id":1,"amount":50.0}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "installment purchase (op=2) with positive amount",
			body:       `{"account_id":1,"operation_type_id":2,"amount":23.5}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "withdrawal (op=3) with positive amount",
			body:       `{"account_id":1,"operation_type_id":3,"amount":18.7}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "credit voucher (op=4) with negative amount",
			body:       `{"account_id":1,"operation_type_id":4,"amount":-60.0}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "missing all fields",
			body:       `{}`,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Post(testServer.URL+"/transactions", "application/json", bytes.NewBufferString(tc.body))
			require.NoError(t, err)
			defer resp.Body.Close()
			require.Equal(t, tc.wantStatus, resp.StatusCode)
		})
	}
}
