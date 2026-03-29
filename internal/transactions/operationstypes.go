package transactions

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
