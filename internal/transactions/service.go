package transactions

import (
	"time"
	"transactions-service/domain"
)

//go:generate mockgen -source=./service.go -destination=./service_mocks_test.go -package=transactions
type Service interface {
	CreateTransaction(transactionRequest Transaction) error
}

type service struct {
	repository Repository
}

var _ Service = &service{}

func NewService(repository Repository) (Service, error) {

	if repository == nil {
		return nil, domain.ErrNoRepo
	}

	return &service{
		repository: repository,
	}, nil
}

// CreateTransaction pass the request to repo layer to save the new transaction
func (s *service) CreateTransaction(transaction Transaction) error {
	transaction.EventDate = time.Now().UTC()
	return s.repository.Save(transaction)
}
