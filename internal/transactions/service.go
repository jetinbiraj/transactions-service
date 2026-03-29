package transactions

import (
	"time"
	"transactions-service/domain"
)

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

func (s *service) CreateTransaction(transaction Transaction) error {
	transaction.EventDate = time.Now().UTC()
	return s.repository.Save(transaction)
}
