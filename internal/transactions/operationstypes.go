package transactions

// OperationType represents all supported operation types for the transaction
type OperationType int

const (
	Purchase OperationType = iota
	InstallmentPurchase
	Withdrawal
	CreditVoucher
)

var OperationTypeMap = map[int]OperationType{
	1: Purchase,
	2: InstallmentPurchase,
	3: Withdrawal,
	4: CreditVoucher,
}

var OperationId = map[OperationType]int{
	Purchase:            1,
	InstallmentPurchase: 2,
	Withdrawal:          3,
	CreditVoucher:       4,
}
