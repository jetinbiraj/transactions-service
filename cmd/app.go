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

	if err := server.RegisterAndServeRouter(ApplicationServer()); err != nil {
		return err
	}

	log.Printf("Set up took %vms", time.Since(startTime))
	return nil
}

func ApplicationServer() server.Server {
	return server.NewServer(config.ServerConfig(),
		accounts.NewHandler(accounts.NewService()),
		transactions.NewHandler(transactions.NewService()))
}
