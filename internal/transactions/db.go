package transactions

type Repository interface {
	CreateTransaction()
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

var _ Repository = &repository{}

func (r *repository) CreateTransaction() {

}
