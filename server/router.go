package server

import (
	"log"
	"net/http"
	"transactions-service/internal/accounts"
	"transactions-service/internal/transactions"
)

func RegisterAndServeRouter(server Server) error {

	mux := http.NewServeMux()

	accounts.RegisterRoutes(mux, server.accountsHandler)
	transactions.RegisterRoutes(mux, server.transactionsHandler)

	log.Printf("Application server staring on port %v", server.config.Port)

	return http.ListenAndServe(server.config.Port, mux)
}
