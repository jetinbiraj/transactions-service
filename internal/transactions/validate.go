package transactions

import "errors"

func validateTransaction(req CreateTransactionRequest) error {
	if req.AccountId <= 0 {
		return errors.New("account_id must be positive")
	}

	if req.OperationTypeId < 1 || req.OperationTypeId > 4 {
		return errors.New("invalid operation_type_id")
	}

	// TODO: Assumption made that transaction amount cannot be 0, change the logic if that's not true
	if req.Amount == 0 {
		return errors.New("amount must not be zero")
	}

	switch OperationTypeMap[req.OperationTypeId] {
	case Purchase, InstallmentPurchase, Withdrawal:
		if req.Amount > 0 {
			return errors.New("amount must be negative for this operation type")
		}
	case CreditVoucher:
		if req.Amount < 0 {
			return errors.New("amount must be positive for credit voucher")
		}
	}

	return nil
}
