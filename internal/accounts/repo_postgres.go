package accounts

import (
	"database/sql"
	"errors"
	"transactions-service/domain"
)

type postgresStore struct {
	db *sql.DB
}

var _ Repository = &postgresStore{}

func NewPostgresStore(db *sql.DB) Repository {
	return &postgresStore{db: db}
}

func (r *postgresStore) Save(account Account) (int64, error) {
	const query = `INSERT INTO accounts (document_number) VALUES ($1) RETURNING account_id`

	var accountID int64
	if err := r.db.QueryRow(query, account.DocumentNumber).Scan(&accountID); err != nil {
		return 0, err
	}

	return accountID, nil
}

func (r *postgresStore) GetById(accountId int64) (*Account, error) {
	const query = `SELECT document_number FROM accounts WHERE account_id = $1`

	var account Account
	err := r.db.QueryRow(query, accountId).Scan(&account.DocumentNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &account, nil
}
