package entities

import (
	"github.com/andrew-waters/frdm/services/locale/core/errors"
	"github.com/andrew-waters/frdm/services/locale/service"
	"github.com/andrew-waters/frdm/services/locale/service/requests"
	"github.com/andrew-waters/frdm/services/locale/service/responses"
)

// Country represents a single country
type Country struct {
	ISO       string `json:"iso"`
	ISO3      string `json:"iso3"`
	NumCode   int    `json:"numcode"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Continent string `json:"continent"`
}

// Response returns a responses.Country for a Country
func (c Country) Response() responses.Country {
	return responses.Country{
		ISO:       c.ISO,
		ISO3:      c.ISO3,
		NumCode:   c.NumCode,
		Name:      c.Name,
		Slug:      c.Slug,
		Continent: c.Continent,
	}
}

// Validate an incoming service request
func (c Country) Validate(request service.Request) error {
	switch r := request.(type) {
	case requests.FindCountriesInContinent:
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
