package memory

import (
	"encoding/json"

	"github.com/andrew-waters/frdm/services/locale/core/entities"
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

func (c CountryRepository) FindCountriesInContinent(ISO string) ([]entities.Country, error) {

	var matching []entities.Country

	for _, country := range c.countries {
		if country.Continent == ISO {
			matching = append(matching, country)
		}
	}

	return matching, nil
}
