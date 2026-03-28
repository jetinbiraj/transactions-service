package server

import (
	"transactions-service/internal/accounts"
	"transactions-service/internal/transactions"
)

type Config struct {
	Port string
}

type Server struct {
	config              Config
	accountsHandler     accounts.Handler
	transactionsHandler transactions.Handler
}

func NewServer(config Config, accountsHandler accounts.Handler, transactionsHandler transactions.Handler) Server {
	return Server{
		config:              config,
		accountsHandler:     accountsHandler,
		transactionsHandler: transactionsHandler,
	}
}
