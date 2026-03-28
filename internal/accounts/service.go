package accounts

type Service interface {
	CreateAccount(accountRequest AccountRequest) error
	GetAccount(accountId string) (AccountInformation, error)
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

func (s *service) CreateAccount(accountRequest AccountRequest) error {
	return s.repository.CreateAccount(accountRequest)
}

func (s *service) GetAccount(accountId string) (AccountInformation, error) {
	return s.repository.GetAccount(accountId)
}
