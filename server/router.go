package server

import (
	"net/http"
	"time"
	"transactions-service/internal/accounts"
	"transactions-service/internal/transactions"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewHTTPServer(server Server) *http.Server {
	return &http.Server{
		Addr:              server.config.Port,
		Handler:           newRouter(server),
		ReadHeaderTimeout: 5 * time.Second,
	}
}

func newRouter(server Server) *http.ServeMux {

	mux := http.NewServeMux()

	accounts.RegisterRoutes(mux, server.accountsHandler)
	transactions.RegisterRoutes(mux, server.transactionsHandler)

	health(mux)
	swagger(mux)

	return mux
}

func health(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}

func swagger(mux *http.ServeMux) {
	mux.HandleFunc("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
}
