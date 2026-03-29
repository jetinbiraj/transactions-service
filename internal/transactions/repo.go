package transactions

import (
	"log"
	"sync"
	"time"
)

// Satisfy the Repository interface in order to use any database

//go:generate mockgen -source=./repo.go -destination=./repo_mocks_test.go -package=transactions
type Repository interface {
	Save(transaction Transaction) error
}

type memoryStore struct {
	mu                   sync.RWMutex
	currentTransactionId int64
	transactions         map[int64]*Transaction
}

var _ Repository = &memoryStore{}

func NewRepository() Repository {

	testData := loadTestData()
	return &memoryStore{
		currentTransactionId: int64(len(testData)),
		transactions:         testData,
	}
}

func loadTestData() map[int64]*Transaction {
	t, err := time.Parse(time.RFC3339Nano, "2020-01-01T10:32:07.7199222Z")
	if err != nil {
		log.Fatalf("failed to parse time, err: %v", err)
	}
	return map[int64]*Transaction{
		1: {
			AccountId:       1,
			OperationTypeId: InstallmentPurchase,
			Amount:          -50.0,
			EventDate:       t,
		},
		2: {
			AccountId:       1,
			OperationTypeId: InstallmentPurchase,
			Amount:          -23.5,
			EventDate:       t,
		},
		3: {
			AccountId:       1,
			OperationTypeId: InstallmentPurchase,
			Amount:          -18.7,
			EventDate:       t,
		},
		4: {
			AccountId:       1,
			OperationTypeId: CreditVoucher,
			Amount:          60.0,
			EventDate:       t,
		},
	}
}

func (r *memoryStore) Save(transaction Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.currentTransactionId++

	r.transactions[r.currentTransactionId] = &transaction

	return nil
}
