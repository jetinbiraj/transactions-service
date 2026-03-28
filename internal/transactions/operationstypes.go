package transactions

type OperationType int

const (
	Purchase OperationType = iota
	InstallmentPurchase
	Withdrawal
	CreditVoucher
)

var OperationTypeMap = map[int]OperationType{
	0: Purchase,
	1: InstallmentPurchase,
	2: Withdrawal,
	3: CreditVoucher,
}
