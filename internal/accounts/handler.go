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

const pathParamAccountId = "account_id"

func NewHandler(logEnabled bool, service Service) (*Handler, error) {

	if service == nil {
		return nil, domain.ErrNoService
	}

	return &Handler{
		logEnabled: logEnabled,
		service:    service,
	}, nil
}

// CreateAccount handles the HTTP POST request to create a new account.
//
//	@Tags			accounts
//	@Description	Create account
//	@Accept			json
//	@Param			RequestBody	body		CreateAccountRequest	true	"create account request body"
//	@Success		201			{string}	string					"CREATED"
//	@failure		400			{string}	string					"BAD REQUEST, if request body is invalid"
//	@failure		500			{string}	string					"INTERNAL SERVER ERROR, server side failure"
//	@Router			/accounts [post]
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

// GetAccount handles the HTTP GET request to get account information by account_id.
//
//	@Tags			accounts
//	@Description	Get account information
//	@Produce		json
//	@Param			account_id	path		integer	true	"account_id"
//	@Success		200			{object}	AccountInformationResponse
//	@failure		400			{string}	string	"BAD REQUEST, if account_id is invalid"
//	@failure		404			{string}	string	"NOT FOUND, if account does not exist for account_id"
//	@failure		500			{string}	string	"INTERNAL SERVER ERROR, server side failure"
//	@Router			/accounts/{account_id} [get]
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
