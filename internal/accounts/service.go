package accounts

type Service interface {
	Save(accountRequest Account) error
	GetById(accountId int64) (AccountInformationResponse, error)
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

func (s *service) Save(accountRequest Account) error {
	return s.repository.Save(accountRequest)
}

func (s *service) GetById(accountId int64) (AccountInformationResponse, error) {

	account, err := s.repository.GetById(accountId)
	if err != nil {
		return AccountInformationResponse{}, err
	}

	return AccountInformationResponse{
		AccountId:      accountId,
		DocumentNumber: account.DocumentNumber,
	}, nil
}
