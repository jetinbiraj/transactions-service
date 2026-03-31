package integration

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"transactions-service/internal/accounts"
	"transactions-service/internal/transactions"
)

var testServer *httptest.Server

func setupTestServer() error {
	accountsService, err := accounts.NewService(accounts.NewMemoryStore())
	if err != nil {
		return err
	}

	accountsHandler, err := accounts.NewHandler(true, accountsService)
	if err != nil {
		return err
	}

	transactionsService, err := transactions.NewService(transactions.NewMemoryStore())
	if err != nil {
		return err
	}

	transactionsHandler, err := transactions.NewHandler(true, transactionsService)
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /accounts", accountsHandler.CreateAccount)
	mux.HandleFunc("GET /accounts/{accountId}", accountsHandler.GetAccount)
	mux.HandleFunc("POST /transactions", transactionsHandler.CreateTransaction)

	testServer = httptest.NewServer(mux)
	return nil
}

func TestMain(m *testing.M) {
	if err := setupTestServer(); err != nil {
		log.Printf("test server setup failed: %v", err)
		os.Exit(1)
	}

	code := m.Run()

	teardown()

	os.Exit(code)
}

func teardown() {
	if testServer != nil {
		testServer.Close()
	}
}
