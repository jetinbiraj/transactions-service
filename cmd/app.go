package main

import (
	"net/http"
	"transactions-service/config"
	"transactions-service/db"
	"transactions-service/internal/accounts"
	"transactions-service/internal/transactions"
	"transactions-service/server"
)

const (
	dbPostgres = "postgres"
)

func buildHTTPServer() (*http.Server, error) {
	appServer, err := applicationServer()
	if err != nil {
		return nil, err
	}

	return server.NewHTTPServer(appServer), nil
}

func applicationServer() (server.Server, error) {

	var accountsRepository accounts.Repository
	var transactionsRepository transactions.Repository

	if config.GetDBName() == dbPostgres {

		err := db.OpenPostgres(config.GetPostgresConfig())
		if err != nil {
			return server.Server{}, err
		}
		accountsRepository = accounts.NewPostgresStore(db.Pg)
		transactionsRepository = transactions.NewPostgresStore(db.Pg)

		db.Init()

	} else {
		accountsRepository = accounts.NewMemoryStore()
		transactionsRepository = transactions.NewMemoryStore()
	}

	accountService, err := accounts.NewService(accountsRepository)
	if err != nil {
		return server.Server{}, err
	}
	transactionService, err := transactions.NewService(transactionsRepository)
	if err != nil {
		return server.Server{}, err
	}

	logEnabled := config.IsLogEnabled()

	accountsHandler, err := accounts.NewHandler(logEnabled, accountService)
	if err != nil {
		return server.Server{}, err
	}

	transactionsHandler, err := transactions.NewHandler(logEnabled, transactionService)
	if err != nil {
		return server.Server{}, err
	}

	return server.NewServer(config.ServerConfig(),
			accountsHandler,
			transactionsHandler,
		),
		nil
}
