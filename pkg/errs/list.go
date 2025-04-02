package errs

import "errors"

var (
	ErrValidation = errors.New("invalid input")
)

const (
	ErrCodeNotFound = 1011
)
