package entities

import (
	"github.com/andrew-waters/frdm/services/locale/core/errors"
	"github.com/andrew-waters/frdm/services/locale/service"
	"github.com/andrew-waters/frdm/services/locale/service/requests"
	"github.com/andrew-waters/frdm/services/locale/service/responses"
)

// Continent represents a single continent
type Continent struct {
	ISO  string `json:"iso"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// Response returns a responses.Continent for a Continent
func (c Continent) Response() responses.Continent {
	return responses.Continent{
		ISO:  c.ISO,
		Name: c.Name,
		Slug: c.Slug,
	}
}

// Validate an incoming service request
func (c Continent) Validate(request service.Request) error {
	switch r := request.(type) {
	case requests.FindContinentByISO:
		isoLength := len(r.ISO)
		if isoLength == 0 {
			return errors.ErrNoISO
		}
		if isoLength < 2 {
			return errors.ErrISOTooShort
		}
		if isoLength > 2 {
			return errors.ErrISOTooLong
		}
		return nil
	}
	return nil
}
