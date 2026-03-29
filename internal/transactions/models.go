package transactions

import "time"

type CreateTransactionRequest struct {
	AccountId       int64   `json:"account_id"`
	OperationTypeId int     `json:"operation_type_id"`
	Amount          float64 `json:"amount"`
}

type Transaction struct {
	AccountId       int64
	OperationTypeId OperationType
	Amount          float64
	EventDate       time.Time
}

//func (r CreateTransactionRequest) TransactionFromCreateTransactionRequest() Transaction {
//	return Transaction{
//		AccountId:       r.AccountId,
//		OperationTypeId: OperationTypeMap[r.OperationTypeId],
//		Amount:          r.Amount,
//	}
//}
