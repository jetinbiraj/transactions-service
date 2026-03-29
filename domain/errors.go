package domain

import "errors"

var (
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrNotFound           = errors.New("not found")
	ErrNoService          = errors.New("no service provided")
	ErrNoRepo             = errors.New("no repo provided")
)
