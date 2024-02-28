package common

import (
	"errors"
)

var (
	ErrTicketNotFound      = errors.New("ticket does not exist")
	ErrCustomfieldNotFound = errors.New("customfield does not exist")
	ErrUnprocessable       = errors.New("invalid parameters")
	ErrRecordNotFound      = errors.New("record not found")
)
