package accounts

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"transactions-service/domain"
	"transactions-service/internal/utils/api"
)

type Handler struct {
	accountId  string
	logEnabled bool
	service    Service
}

func NewHandler(logEnabled bool, service Service) (*Handler, error) {

	if service == nil {
		return nil, domain.ErrNoService
	}

	return &Handler{
		accountId:  "accountId",
		logEnabled: logEnabled,
		service:    service,
	}, nil
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		api.Error(w, r, domain.ErrInvalidRequestBody, 0, h.logEnabled)
		return
	}
	defer r.Body.Close()

	if err := validateDocumentNumber(reqBody.DocumentNumber); err != nil {
		api.Error(w, r, err, http.StatusBadRequest, h.logEnabled)
		return
	}

	if err := h.service.CreateAccount(reqBody.AccountFromCreateAccountRequest()); err != nil {
		api.Error(w, r, err, 0, h.logEnabled)
		return
	}

	api.SuccessJson(w, r, domain.MessageResponse{Message: "account created successfully"},
		http.StatusCreated, h.logEnabled)
}

func (h *Handler) GetAccount(w http.ResponseWriter, r *http.Request) {
	accountId := r.PathValue(h.accountId)
	if err := validateAccountId(accountId); err != nil {
		api.Error(w, r, err, http.StatusBadRequest, h.logEnabled)
		return
	}
	id, err := strconv.ParseInt(accountId, 10, 64)
	if err != nil || id <= 0 {
		api.Error(w, r, errors.New("invalid account_id"), http.StatusBadRequest, h.logEnabled)
		return
	}

	accountInfo, err := h.service.GetAccount(id)
	if err != nil {
		api.Error(w, r, err, 0, h.logEnabled)
		return
	}

	api.SuccessJson(w, r, accountInfo, 0, h.logEnabled)
}
