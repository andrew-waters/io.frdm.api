package interactors

import (
	"github.com/andrew-waters/frdm/services/locale/core/entities"
	"github.com/andrew-waters/frdm/services/locale/core/repositories"
	"github.com/andrew-waters/frdm/services/locale/service/requests"
	"github.com/andrew-waters/frdm/services/locale/service/responses"
)

// CountryInteractor contains the repository for interacting with countries
type CountryInteractor struct {
	repository repositories.Countries
}

// NewCountryInteractor returns a usable CountryInteractor
func NewCountryInteractor(repository repositories.Countries) *CountryInteractor {
	return &CountryInteractor{
		repository: repository,
	}
}

// FindCountryWithISO gets a country by ISO(2)
func (interactor CountryInteractor) FindCountryWithISO(request requests.FindCountryWithISO) (responses.Country, error) {

	var country entities.Country

	// validate the incoming request
	if err := country.Validate(request); err != nil {
		return responses.Country{}, err
	}

	// finished validation, find the continent
	country, err := interactor.repository.FindCountryWithISO(request.ISO)
	if err != nil {
		return responses.Country{}, err
	}

	// return a Page response from the entity
	return country.Response(), nil
}

// FindCountriesInContinent retrieves all countries in a continent
func (interactor CountryInteractor) FindCountriesInContinent(request requests.FindCountriesInContinent) ([]responses.Country, error) {

	var response []responses.Country
	var country entities.Country
	var countries []entities.Country

	// validate the incoming request
	if err := country.Validate(request); err != nil {
		return response, err
	}

	// finished validation, find the countries
	countries, err := interactor.repository.FindCountriesInContinent(request.ISO)
	if err != nil {
		return response, err
	}

	// format and return the response
	for _, country := range countries {
		response = append(response, country.Response())
	}
	return response, nil
}

func (interactor CountryInteractor) FindAllCountries(request requests.FindAllCountries) ([]responses.Country, error) {

	var response []responses.Country
	var country entities.Country

	// validate the incoming request
	if err := country.Validate(request); err != nil {
		return response, err
	}

	// finished validation, find the countries
	countries, err := interactor.repository.FindAllCountries()
	if err != nil {
		return response, err
	}

	// format and return the response
	for _, country := range countries {
		response = append(response, country.Response())
	}
	return response, nil
}
