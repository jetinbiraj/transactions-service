package accounts

type Repository interface {
	CreateAccount(accountRequest AccountRequest) error
	GetAccount(accountId string) (AccountInformation, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

var _ Repository = &repository{}

func (r *repository) CreateAccount(accountRequest AccountRequest) error {

	// TODO: Implement me!
	return nil
}

func (r *repository) GetAccount(accountId string) (AccountInformation, error) {

	// TODO: Implement me!
	return AccountInformation{}, nil
}
