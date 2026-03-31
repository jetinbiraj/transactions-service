package main

import (
	"net/http"
	"transactions-service/config"
	"transactions-service/internal/accounts"
	"transactions-service/internal/transactions"
	"transactions-service/server"
)

func buildHTTPServer() (*http.Server, error) {
	appServer, err := applicationServer()
	if err != nil {
		return nil, err
	}

	return server.NewHTTPServer(appServer), nil
}

func applicationServer() (server.Server, error) {

	logEnabled := config.IsLogEnabled()

	accountsHandler, err := getAccountsHandler(logEnabled)
	if err != nil {
		return server.Server{}, err
	}

	transactionsHandler, err := getTransactionsHandler(logEnabled)
	if err != nil {
		return server.Server{}, err
	}

	return server.NewServer(config.ServerConfig(),
			accountsHandler,
			transactionsHandler,
		),
		nil
}

func getAccountsHandler(logEnabled bool) (*accounts.Handler, error) {
	accountService, err := accounts.NewService(accounts.NewMemoryStore())
	if err != nil {
		return nil, err
	}

	accountsHandler, err := accounts.NewHandler(logEnabled, accountService)
	if err != nil {
		return nil, err
	}

	return accountsHandler, nil
}

func getTransactionsHandler(logEnabled bool) (*transactions.Handler, error) {
	transactionService, err := transactions.NewService(transactions.NewMemoryStore())
	if err != nil {
		return nil, err
	}

	transactionsHandler, err := transactions.NewHandler(logEnabled, transactionService)
	if err != nil {
		return nil, err
	}

	return transactionsHandler, nil
}
