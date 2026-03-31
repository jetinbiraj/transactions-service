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
	db := config.GetDBName()

	accountsHandler, err := getAccountsHandler(logEnabled, db)
	if err != nil {
		return server.Server{}, err
	}

	transactionsHandler, err := getTransactionsHandler(logEnabled, db)
	if err != nil {
		return server.Server{}, err
	}

	return server.NewServer(config.ServerConfig(),
			accountsHandler,
			transactionsHandler,
		),
		nil
}

func getAccountsHandler(logEnabled bool, db string) (*accounts.Handler, error) {

	var repository accounts.Repository
	if db == "postgres" {
		// TODO: Add postgres as persistent storage
	} else {
		repository = accounts.NewMemoryStore()
	}

	accountService, err := accounts.NewService(repository)
	if err != nil {
		return nil, err
	}

	accountsHandler, err := accounts.NewHandler(logEnabled, accountService)
	if err != nil {
		return nil, err
	}

	return accountsHandler, nil
}

func getTransactionsHandler(logEnabled bool, db string) (*transactions.Handler, error) {
	var repository transactions.Repository
	if db == "postgres" {
		// TODO: Add postgres as persistent storage
	} else {
		repository = transactions.NewMemoryStore()
	}

	transactionService, err := transactions.NewService(repository)
	if err != nil {
		return nil, err
	}

	transactionsHandler, err := transactions.NewHandler(logEnabled, transactionService)
	if err != nil {
		return nil, err
	}

	return transactionsHandler, nil
}
