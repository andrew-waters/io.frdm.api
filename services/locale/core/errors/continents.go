package errors

import "errors"

var (
	// ErrContinentNotFound is returned when no ISO can be found
	ErrContinentNotFound = errors.New("no continent found with the supplied ISO")
	// ErrNoISO is returned when no ISO is supplied
	ErrNoISO = errors.New("no ISO supplied")
	// ErrISOTooLong is returned when a supplied ISO is too long
	ErrISOTooLong = errors.New("supplied ISO is too long")
	// ErrISOTooShort is returned when a supplied ISO is too short
	ErrISOTooShort = errors.New("supplied ISO is too short")
)
