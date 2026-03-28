package transactions

type Service interface {
	CreateTransaction()
}

type service struct {
}

var _ Service = &service{}

func NewService() Service {
	return &service{}
}

func (s *service) CreateTransaction() {

}
