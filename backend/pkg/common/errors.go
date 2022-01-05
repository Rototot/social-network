package common

import "errors"

var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
	ErrInvalidJson   = errors.New("http: invalid json")
)
