package memory

import (
	"encoding/json"

	"github.com/andrew-waters/frdm/services/locale/core/entities"
	"github.com/andrew-waters/frdm/services/locale/core/errors"
)

type CountryRepository struct {
	countries []entities.Country
}

func NewCountryRepository() (*CountryRepository, error) {

	// read all continents from disk and add them to memory
	var countries []entities.Country
	fileData, err := fileContents("./datastore/memory/countries.json")
	if err != nil {
		return &CountryRepository{}, err
	}
	json.Unmarshal(fileData, &countries)

	// return the repository
	return &CountryRepository{
		countries: countries,
	}, nil
}

func (c CountryRepository) FindCountryWithISO(ISO string) (entities.Country, error) {

	for _, country := range c.countries {
		if country.ISO == ISO {
			return country, nil
		}
	}

	return entities.Country{}, errors.ErrCountryNotFound
}

func (c CountryRepository) FindCountriesInContinent(ISO string) ([]entities.Country, error) {

	var matching []entities.Country
	for _, country := range c.countries {
		if country.Continent == ISO {
			matching = append(matching, country)
		}
	}

	return matching, nil
}

func (c CountryRepository) FindAllCountries() ([]entities.Country, error) {

	var matching []entities.Country
	for _, country := range c.countries {
		matching = append(matching, country)
	}

	return matching, nil

}
