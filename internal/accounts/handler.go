package accounts

type Handler interface {
}

type handler struct {
}

var _ Handler = &handler{}

func NewHandler() Handler {
	return &handler{}
}
