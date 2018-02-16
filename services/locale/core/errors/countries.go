package errors

import "errors"

var (
	// ErrCountryNotFound is returned when a country cannot be found
	ErrCountryNotFound = errors.New("Country not found")
)
