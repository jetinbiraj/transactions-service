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

// CreateAccount handles the HTTP POST request to create a new account.
//
//	@Tags			accounts
//	@Description	Create account
//	@Accept			json
//	@Param			RequestBody	body		CreateAccountRequest	true	"create account request body"
//	@Success		201			{object}	AccountInformationResponse
//	@failure		400			{object}	domain.ErrorResponse	"BAD REQUEST, if request body is invalid"
//	@failure		500			{object}	domain.ErrorResponse	"INTERNAL SERVER ERROR, server side failure"
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

	createdAccountInformation, err := h.service.CreateAccount(Account{
		DocumentNumber: reqBody.DocumentNumber,
	})
	if err != nil {
		api.Error(w, r, err, 0, h.logEnabled)
		return
	}

	api.SuccessJson(w, r, createdAccountInformation,
		http.StatusCreated, h.logEnabled)
}

// GetAccount handles the HTTP GET request to get account information by accountId.
//
//	@Tags			accounts
//	@Description	Get account information
//	@Produce		json
//	@Param			accountId	path		integer	true	"accountId"
//	@Success		200			{object}	AccountInformationResponse
//	@failure		400			{object}	domain.ErrorResponse	"BAD REQUEST, if accountId is invalid"
//	@failure		404			{object}	domain.ErrorResponse	"NOT FOUND, if account does not exist for accountId"
//	@failure		500			{object}	domain.ErrorResponse	"INTERNAL SERVER ERROR, server side failure"
//	@Router			/accounts/{accountId} [get]
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
