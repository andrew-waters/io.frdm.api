package repositories

import "github.com/andrew-waters/frdm/services/locale/core/entities"

// Countries provides access to country storage
type Countries interface {
	FindCountryWithISO(ISO string) (entities.Country, error)
	FindCountriesInContinent(ISO string) ([]entities.Country, error)
	FindAllCountries() ([]entities.Country, error)
}
