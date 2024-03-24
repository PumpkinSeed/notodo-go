package storage

import "errors"

var (
	ErrNotFound            = errors.New("item not found in storage")
	ErrUnknownStorageError = errors.New("unknown storage error")
)
