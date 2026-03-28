package transactions

type Service interface {
	CreateTransaction(transactionRequest TransactionRequest) error
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

func (s *service) CreateTransaction(transactionRequest TransactionRequest) error {

	// TODO: Implement me!
	return nil
}
