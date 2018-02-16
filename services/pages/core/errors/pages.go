package errors

import "errors"

var (
	// ErrClashingURI is returned when there is a URI clash
	ErrClashingURI = errors.New("page with this URI already exists")
)

// return responses.Page{}, errors.ErrClashingURI
