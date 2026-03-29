package accounts

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"transactions-service/domain"
	"transactions-service/internal/utils/api"
)

type Handler interface {
	CreateAccount(w http.ResponseWriter, r *http.Request)
	GetAccount(w http.ResponseWriter, r *http.Request)
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

func (h *handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var reqBody CreateAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		api.Error(w, r, domain.ErrInvalidRequestBody, 0)
		return
	}
	defer r.Body.Close()

	if err := validateDocumentNumber(reqBody.DocumentNumber); err != nil {
		api.Error(w, r, err, http.StatusBadRequest)
		return
	}

	if err := h.service.Save(reqBody.AccountFromCreateAccountRequest()); err != nil {
		api.Error(w, r, err, 0)
		return
	}

	api.SuccessJsonWithStatusCode(w, r, domain.MessageResponse{Message: "account created successfully"}, http.StatusCreated)
}

func (h *handler) GetAccount(w http.ResponseWriter, r *http.Request) {
	accountId := r.PathValue("accountId")
	if err := validateAccountId(accountId); err != nil {
		api.Error(w, r, err, http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseInt(accountId, 10, 64)
	if err != nil || id <= 0 {
		api.Error(w, r, errors.New("invalid account_id"), http.StatusBadRequest)
		return
	}

	accountInfo, err := h.service.GetById(id)
	if err != nil {
		api.Error(w, r, err, 0)
		return
	}

	api.SuccessJson(w, r, accountInfo)
}
