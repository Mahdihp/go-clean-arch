package domain

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal WsServer Error happen
	ErrInternalServerError = errors.New("internal WsServer Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("your requested Item is not found")
	// ErrConflict will throw if the current action already exists
	ErrConflict = errors.New("your Item already exist")
	// ErrBadParamInput will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("given Param is not valid")
)
