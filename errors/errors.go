package errors

import "errors"

type Error error

var (
	ErrStoreTypeNotSupported = errors.New("store type not supported")
	ErrInvalidHexValue       = errors.New("invalid hex value")
)
