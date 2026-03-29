package accounts

import (
	"encoding/json"
	"net/http"
	"strconv"
	"transactions-service/domain"
	"transactions-service/internal/utils/api"
)

type Handler struct {
	logEnabled bool
	service    Service
}

const pathParamAccountId = "accountId"

func NewHandler(logEnabled bool, service Service) (*Handler, error) {

	if service == nil {
		return nil, domain.ErrNoService
	}

	return &Handler{
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

	if err := h.service.CreateAccount(Account{
		DocumentNumber: reqBody.DocumentNumber,
	}); err != nil {
		api.Error(w, r, err, 0, h.logEnabled)
		return
	}

	api.SuccessJson(w, r, domain.MessageResponse{Message: "account created successfully"},
		http.StatusCreated, h.logEnabled)
}

func (h *Handler) GetAccount(w http.ResponseWriter, r *http.Request) {
	accountId := r.PathValue(pathParamAccountId)
	if err := validateAccountId(accountId); err != nil {
		api.Error(w, r, err, http.StatusBadRequest, h.logEnabled)
		return
	}

	id, _ := strconv.ParseInt(accountId, 10, 64)
	accountInfo, err := h.service.GetAccount(id)
	if err != nil {
		api.Error(w, r, err, 0, h.logEnabled)
		return
	}

	api.SuccessJson(w, r, accountInfo, 0, h.logEnabled)
}
