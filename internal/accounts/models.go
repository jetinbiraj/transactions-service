package accounts

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

type GetAccountResponse struct {
	AccountId      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}
