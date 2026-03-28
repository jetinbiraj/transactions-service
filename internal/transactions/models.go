package transactions

type CreateTransactionRequest struct {
	AccountId       int     `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

type TransactionRequest struct {
	AccountId       int           `json:"account_id"`
	OperationTypeId OperationType `json:"operation_type_id"`
	Amount          float64       `json:"amount"`
}

func (r CreateTransactionRequest) GetTransactionRequestFromCreateTransactionRequest() TransactionRequest {
	return TransactionRequest{
		AccountId:       r.AccountId,
		OperationTypeId: OperationTypeMap[r.OperationTypeId],
		Amount:          r.Amount,
	}
}
