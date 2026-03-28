package domain

import "errors"

var (
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrNotFound           = errors.New("not found")
)
