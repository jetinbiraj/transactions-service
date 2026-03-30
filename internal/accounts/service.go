package accounts

import "transactions-service/domain"

//go:generate mockgen -source=./service.go -destination=./service_mocks_test.go -package=accounts
type Service interface {
	CreateAccount(accountRequest Account) error
	GetAccount(accountId int64) (AccountInformationResponse, error)
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

// CreateAccount creates a new account for a given document number.
func (s *service) CreateAccount(accountRequest Account) error {
	return s.repository.Save(accountRequest)
}

// GetAccount gets the Account associated with provided accountId.
func (s *service) GetAccount(accountId int64) (AccountInformationResponse, error) {

	account, err := s.repository.GetById(accountId)
	if err != nil {
		return AccountInformationResponse{}, err
	}

	return AccountInformationResponse{
		AccountId:      accountId,
		DocumentNumber: account.DocumentNumber,
	}, nil
}
