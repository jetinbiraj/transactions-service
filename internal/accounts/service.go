package accounts

type Service interface {
	CreateAccount()
	GetAccount()
}

type service struct {
}

var _ Service = &service{}

func NewService() Service {
	return &service{}
}

func (s *service) CreateAccount() {

}

func (s *service) GetAccount() {

}
