package service

import (
	"github.com/andrew-waters/frdm/services/locale/service/requests"
	"github.com/andrew-waters/frdm/services/locale/service/responses"
)

// Countries is a boundary that can do things with countries
type Countries interface {
	FindCountryWithISO(requests.FindCountryWithISO) (responses.Country, error)
	FindCountriesInContinent(requests.FindCountriesInContinent) ([]responses.Country, error)
	FindAllCountries(requests.FindAllCountries) ([]responses.Country, error)
}
