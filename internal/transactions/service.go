package transactions

import "time"

type Service interface {
	CreateTransaction(transactionRequest Transaction) error
}

type service struct {
	repository Repository
}

var _ Service = &service{}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateTransaction(transaction Transaction) error {
	transaction.EventDate = time.Now().UTC()
	return s.repository.Save(transaction)
}
