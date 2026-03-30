package accounts

// CreateAccountRequest represents the request body for create account request
type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number" format:"long" minLength:"11" maxLength:"11" example:"12345678900"`
}

// AccountInformationResponse represents the response body for get account
type AccountInformationResponse struct {
	AccountId      int64  `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

// Account represents data transfer object in service and db entity for accounts
type Account struct {
	DocumentNumber string
}
