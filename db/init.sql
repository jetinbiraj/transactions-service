-- USER postgres;
--
-- DROP
--     DATABASE IF EXISTS transactionservice;
--
-- CREATE
--     DATABASE transactionservice;

CREATE TABLE accounts
(
    account_id      BIGSERIAL PRIMARY KEY,
    document_number VARCHAR(11) NOT NULL
);

INSERT INTO accounts (document_number)
VALUES ('12345678900');

CREATE TABLE transactions
(
    transaction_id    BIGSERIAL PRIMARY KEY,
    account_id        BIGINT         NOT NULL,
    operation_type_id INT            NOT NULL,
    amount            NUMERIC(10, 2) NOT NULL,
    event_date        TIMESTAMP      NOT NULL DEFAULT NOW()
);

INSERT INTO transactions (account_id, operation_type_id, amount, event_date)
VALUES (1, 1, -50.0, '2020-01-01T10:32:07.7199222');
INSERT INTO transactions (account_id, operation_type_id, amount, event_date)
VALUES (1, 1, -23.5, '2020-01-01T10:32:07.7199222');
INSERT INTO transactions (account_id, operation_type_id, amount, event_date)
VALUES (1, 1, -18.7, '2020-01-01T10:32:07.7199222');
INSERT INTO transactions (account_id, operation_type_id, amount, event_date)
VALUES (1, 4, -50.0, '2020-01-01T10:32:07.7199222');

-- CREATE UNIQUE INDEX idx_accounts_document ON accounts (document_number);
-- CREATE INDEX idx_transactions_account ON transactions (account_id);