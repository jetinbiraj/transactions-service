package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"transactions-service/internal/accounts"

	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// POST /accounts
// ---------------------------------------------------------------------------

func TestAccountCreation_InvalidCases(t *testing.T) {
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
			name:       "empty document_number",
			body:       `{"document_number":""}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "document_number too short (10 digits)",
			body:       `{"document_number":"1234567890"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "document_number too long (12 digits)",
			body:       `{"document_number":"123456789012"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "document_number contains letters",
			body:       `{"document_number":"1234567890a"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "document_number contains special characters",
			body:       `{"document_number":"1234567890!"}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "missing document_number field",
			body:       `{}`,
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Post(testServer.URL+"/accounts", "application/json", bytes.NewBufferString(tc.body))
			require.NoError(t, err)
			defer resp.Body.Close()
			require.Equal(t, tc.wantStatus, resp.StatusCode)
		})
	}
}

// ---------------------------------------------------------------------------
// GET /accounts/{accountId}
// ---------------------------------------------------------------------------

func TestGetAccount_Success(t *testing.T) {
	// account id=1 is preloaded in the memory store
	resp, err := http.Get(fmt.Sprintf("%s/accounts/1", testServer.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode)

	var body accounts.AccountInformationResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&body))
	require.Equal(t, int64(1), body.AccountId)
	require.NotEmpty(t, body.DocumentNumber)
}

func TestGetAccount_NotFound(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("%s/accounts/999999", testServer.URL))
	require.NoError(t, err)
	defer resp.Body.Close()

	require.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestGetAccount_InvalidCases(t *testing.T) {
	tests := []struct {
		name       string
		accountId  string
		wantStatus int
	}{
		{
			name:       "non-numeric account_id",
			accountId:  "abc",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "zero account_id",
			accountId:  "0",
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "negative account_id",
			accountId:  "-1",
			wantStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("%s/accounts/%s", testServer.URL, tc.accountId))
			require.NoError(t, err)
			defer resp.Body.Close()
			require.Equal(t, tc.wantStatus, resp.StatusCode)
		})
	}
}

func TestAccountCreation_Success(t *testing.T) {

	jsonReq, err := json.Marshal(accounts.CreateAccountRequest{DocumentNumber: "12345678900"})
	require.NoError(t, err)

	resp, err := http.Post(testServer.URL+"/accounts", "application/json", bytes.NewReader(jsonReq))
	if err != nil {
		t.Fatalf("failed to create account: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}

	var accResp accounts.AccountInformationResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&accResp))

	resp, err = http.Get(fmt.Sprintf("%s/accounts/%d", testServer.URL, accResp.AccountId))
	if err != nil {
		t.Fatalf("failed to get account: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var getResp accounts.AccountInformationResponse
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&getResp))

	if getResp.AccountId != accResp.AccountId || accResp.DocumentNumber != getResp.DocumentNumber {
		t.Errorf("account mismatch")
	}
}
