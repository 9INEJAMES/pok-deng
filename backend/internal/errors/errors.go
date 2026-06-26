package errors

import "errors"

var (
	ErrSessionNotFound  = errors.New("ERR_SESSION_NOT_FOUND")
	ErrInvalidState     = errors.New("ERR_INVALID_STATE")
	ErrInvalidAmount    = errors.New("ERR_INVALID_AMOUNT")
	ErrInsufficientFund = errors.New("ERR_INSUFFICIENT_FUNDS")
	ErrInvalidAction    = errors.New("ERR_INVALID_ACTION")
	ErrInvalidRequest   = errors.New("ERR_INVALID_REQUEST")
)
