package main

import (
	"log"
	"time"
	"transactions-service/config"
	"transactions-service/internal/accounts"
	"transactions-service/internal/transactions"
	"transactions-service/server"
)

func StartTransactionsService() error {
	startTime := time.Now()
	if err := config.Set(); err != nil {
		return err
	}

	appServer, err := ApplicationServer()
	if err != nil {
		return err
	}

	if err = server.RegisterAndServeRouter(appServer); err != nil {
		return err
	}

	log.Printf("Set up took %vms", time.Since(startTime))
	return nil
}

func ApplicationServer() (server.Server, error) {

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
	accountService, err := accounts.NewService(accounts.NewRepository())
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
	transactionService, err := transactions.NewService(transactions.NewRepository())
	if err != nil {
		return nil, err
	}

	transactionsHandler, err := transactions.NewHandler(logEnabled, transactionService)
	if err != nil {
		return nil, err
	}

	return transactionsHandler, nil
}
