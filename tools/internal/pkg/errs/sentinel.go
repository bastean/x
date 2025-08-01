package errs

import (
	"errors"
)

var (
	ErrRequiredFlags = errors.New("please define required flags")
)
