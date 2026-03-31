package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

var Pg *sql.DB

func OpenPostgres(cfg PostgresConfig) error {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)

	var err error

	Pg, err = sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	if cfg.MaxOpenConns > 0 {
		Pg.SetMaxOpenConns(cfg.MaxOpenConns)
	}
	if cfg.MaxIdleConns > 0 {
		Pg.SetMaxIdleConns(cfg.MaxIdleConns)
	}
	if cfg.ConnMaxLifetime > 0 {
		Pg.SetConnMaxLifetime(cfg.ConnMaxLifetime)
	}

	if err = Pg.Ping(); err != nil {
		_ = Pg.Close()
		return err
	}
	return nil
}

func Init() {
	initAccounts()
	initTransactions()
}

func initAccounts() {
	_, err := Pg.Exec(`
	CREATE TABLE IF NOT EXISTS accounts (
		account_id SERIAL PRIMARY KEY,
		document_number TEXT NOT NULL
	);`)

	if err != nil {
		log.Printf("error creating accounts: %v", err)
	}
}

func initTransactions() {
	_, err := Pg.Exec(`
CREATE TABLE IF NOT EXISTS transactions
(
    transaction_id    BIGSERIAL PRIMARY KEY,
    account_id        BIGINT         NOT NULL,
    operation_type_id INT            NOT NULL,
    amount            NUMERIC(10, 2) NOT NULL,
    event_date        TIMESTAMP      NOT NULL DEFAULT NOW()
);`)

	if err != nil {
		log.Printf("error creating transactions: %v", err)
	}
}
