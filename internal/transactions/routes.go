package transactions

import "net/http"

func RegisterRoutes(mux *http.ServeMux, handler *Handler) {
	mux.HandleFunc("POST /transactions", handler.CreateTransaction)
}
