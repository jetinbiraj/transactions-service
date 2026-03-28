package transactions

import (
	"encoding/json"
	"net/http"
	"transactions-service/domain"
	"transactions-service/internal/utils/api"
)

type Handler interface {
	CreateTransaction(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service Service
}

var _ Handler = &handler{}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func (h handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		api.Error(w, r, domain.ErrInvalidRequestBody, 0)
		return
	}
	defer r.Body.Close()

	if err := validateTransaction(reqBody); err != nil {
		api.Error(w, r, err, http.StatusBadRequest)
		return
	}
	if err := h.service.CreateTransaction(reqBody.GetTransactionRequestFromCreateTransactionRequest()); err != nil {
		api.Error(w, r, err, 0)
		return
	}

	api.SuccessJsonWithStatusCode(w, r, domain.MessageResponse{Message: "transaction created successfully"}, http.StatusCreated)
}
