CREATE TABLE accounts
(
    account_id      BIGSERIAL PRIMARY KEY,
    document_number VARCHAR(11) NOT NULL
);

CREATE TABLE transactions
(
    transaction_id    BIGSERIAL PRIMARY KEY,
    account_id        BIGINT         NOT NULL,
    operation_type_id INT            NOT NULL,
    amount            NUMERIC(10, 2) NOT NULL,
    event_date        TIMESTAMP      NOT NULL DEFAULT NOW(),
    FOREIGN KEY (account_id) REFERENCES accounts (account_id)
);

-- CREATE UNIQUE INDEX idx_accounts_document ON accounts (document_number);
CREATE INDEX idx_transactions_account ON transactions (account_id);