package transactions

import (
	"time"
	"transactions-service/domain"
)

//go:generate mockgen -source=./service.go -destination=./service_mocks_test.go -package=transactions
type Service interface {
	CreateTransaction(transactionRequest Transaction) (TransactionResponse, error)
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

// CreateTransaction pass the request to repo layer to save the new transaction and created returns the transaction response
func (s *service) CreateTransaction(transaction Transaction) (TransactionResponse, error) {
	transaction.EventDate = time.Now().UTC()

	transactionId, err := s.repository.Save(transaction)
	if err != nil {
		return TransactionResponse{}, err
	}

	return TransactionResponse{
		TransactionId:   transactionId,
		AccountId:       transaction.AccountId,
		OperationTypeId: OperationId[transaction.OperationTypeId],
		Amount:          transaction.Amount,
		EventDate:       transaction.EventDate,
	}, nil
}
