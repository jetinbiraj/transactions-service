package accounts

import (
	"sync"
	"transactions-service/domain"
)

// Satisfy the Repository interface in order to use any database

type Repository interface {
	Save(accountRequest Account) error
	GetById(accountId int64) (*Account, error)
}

type memoryStore struct {
	mu               sync.RWMutex
	currentAccountId int64
	accounts         map[int64]*Account
}

var _ Repository = &memoryStore{}

func NewRepository() Repository {

	testData := loadTestData()

	return &memoryStore{
		currentAccountId: int64(len(testData)),
		accounts:         testData,
	}
}

func loadTestData() map[int64]*Account {
	return map[int64]*Account{
		1: {
			DocumentNumber: "12345678900",
		},
	}
}

func (r *memoryStore) Save(account Account) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	r.currentAccountId++

	r.accounts[r.currentAccountId] = &account

	// TODO: Assumption: multiple accounts possible for same document number

	return nil
}

func (r *memoryStore) GetById(accountId int64) (*Account, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if account, ok := r.accounts[accountId]; ok {
		return account, nil
	}
	return nil, domain.ErrNotFound
}
