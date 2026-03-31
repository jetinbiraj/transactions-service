package transactions

import "database/sql"

type postgresStore struct {
	db *sql.DB
}

var _ Repository = &postgresStore{}

func NewPostgresStore(db *sql.DB) Repository {
	return &postgresStore{db: db}
}

func (r *postgresStore) Save(transaction Transaction) (int64, error) {
	const query = `
		INSERT INTO transactions (account_id, operation_type_id, amount, event_date)
		VALUES ($1, $2, $3, $4)
		RETURNING transaction_id`

	var transactionID int64
	err := r.db.QueryRow(
		query,
		transaction.AccountId,
		OperationId[transaction.OperationTypeId],
		transaction.Amount,
		transaction.EventDate,
	).Scan(&transactionID)
	if err != nil {
		return 0, err
	}

	return transactionID, nil
}
