package accounts

import (
	"fmt"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, handler *Handler) {
	mux.HandleFunc("POST /accounts", handler.CreateAccount)
	mux.HandleFunc(fmt.Sprintf("GET /accounts/{%v}", handler.accountId), handler.GetAccount)
}
