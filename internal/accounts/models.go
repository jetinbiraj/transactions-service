package accounts

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

type Account struct {
	DocumentNumber string `json:"document_number"`
}

type AccountInformationResponse struct {
	AccountId      int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}
