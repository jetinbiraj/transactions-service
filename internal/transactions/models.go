package transactions

import "time"

// CreateTransactionRequest represents the request body for create transaction request
type CreateTransactionRequest struct {
	AccountId       int64   `json:"account_id" format:"integer" example:"1"`
	OperationTypeId int     `json:"operation_type_id" format:"integer" enums:"1,2,3,4" example:"4"`
	Amount          float64 `json:"amount" format:"number" example:"123.45"`
}

// TransactionResponse represents the response body for create transaction request
type TransactionResponse struct {
	TransactionId   int64     `json:"transaction_id"`
	AccountId       int64     `json:"account_id"`
	OperationTypeId int       `json:"operation_type_id"`
	Amount          float64   `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}

// Transaction represents data transfer object in service and db entity for transactions
type Transaction struct {
	AccountId       int64
	OperationTypeId OperationType
	Amount          float64
	EventDate       time.Time
}
