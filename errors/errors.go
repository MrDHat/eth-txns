package errors

import "errors"

type Error error

var (
	ErrStoreTypeNotSupported  = errors.New("store type not supported")
	ErrInvalidHexValue        = errors.New("invalid hex value")
	ErrFailedToGetLatestBlock = errors.New("failed to get the latest block")
	ErrAddressNotFound        = errors.New("address not found")
	ErrInvalidCommand         = errors.New("invalid command")
)
