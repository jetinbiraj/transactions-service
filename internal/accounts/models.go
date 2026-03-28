package accounts

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

type AccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

type AccountInformation struct {
	AccountId      int    `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

func (r CreateAccountRequest) GetAccountRequestFromCreateAccountRequest() AccountRequest {
	return AccountRequest{
		DocumentNumber: r.DocumentNumber,
	}
}
