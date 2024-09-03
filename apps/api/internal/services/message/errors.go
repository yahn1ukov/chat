package message

import "errors"

var (
	ErrTextRequired = errors.New("text is required")
	ErrTextLimit    = errors.New("text can't be more than 200 characters")
)
