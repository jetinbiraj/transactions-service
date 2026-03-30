package server

import (
	"log"
	"net/http"
	"transactions-service/internal/accounts"
	"transactions-service/internal/transactions"
	_ "transactions-service/swagger"

	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterAndServeRouter(server Server) error {

	mux := http.NewServeMux()

	accounts.RegisterRoutes(mux, server.accountsHandler)
	transactions.RegisterRoutes(mux, server.transactionsHandler)

	mux.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	log.Printf("Application server staring on port %v", server.config.Port)

	return http.ListenAndServe(server.config.Port, mux)
}
