package transactions

import (
	"encoding/json"
	"net/http"
	"transactions-service/domain"
	"transactions-service/internal/utils/api"
)

type Handler struct {
	logEnabled bool
	service    Service
}

func NewHandler(logEnabled bool, service Service) (*Handler, error) {

	if service == nil {
		return nil, domain.ErrNoService
	}

	return &Handler{
		logEnabled: logEnabled,
		service:    service,
	}, nil
}

func (h Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		api.Error(w, r, domain.ErrInvalidRequestBody, 0, h.logEnabled)
		return
	}
	defer r.Body.Close()

	if err := validateTransaction(reqBody); err != nil {
		api.Error(w, r, err, http.StatusBadRequest, h.logEnabled)
		return
	}
	if err := h.service.CreateTransaction(reqBody.TransactionFromCreateTransactionRequest()); err != nil {
		api.Error(w, r, err, 0, h.logEnabled)
		return
	}

	api.SuccessJson(w, r, domain.MessageResponse{Message: "transaction created successfully"},
		http.StatusCreated, h.logEnabled)
}
