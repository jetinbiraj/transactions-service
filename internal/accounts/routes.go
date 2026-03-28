package accounts

import "net/http"

func RegisterRoutes(mux *http.ServeMux, handler Handler) {
	mux.HandleFunc("POST /accounts", handler.CreateAccount)
	mux.HandleFunc("GET /accounts/{accountId}", handler.GetAccount)
}
