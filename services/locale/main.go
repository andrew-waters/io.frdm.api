package main

import (
	"os"

	"github.com/andrew-waters/frdm/services/locale/core/interactors"
	"github.com/andrew-waters/frdm/services/locale/datastore/memory"
	"github.com/andrew-waters/frdm/services/locale/service"
	"github.com/andrew-waters/frdm/services/locale/service/requests"
	"github.com/davecgh/go-spew/spew"
)

func main() {

	// create concrete repositories
	var errs []error
	continentRepository, err := memory.NewContinentRepository()
	if err != nil {
		errs = append(errs, err)
	}
	countryRepository, err := memory.NewCountryRepository()
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		spew.Dump(errs)
		os.Exit(1)
	}

	// create the service API
	frdm := NewAPI(
		interactors.NewContinentInteractor(continentRepository),
		interactors.NewCountryInteractor(countryRepository),
	)

	continentISO := "SA"
	countryISO := "GB"

	continent, err := frdm.Continents.FindByISO(requests.FindContinentByISO{ISO: continentISO})
	if err != nil {
		spew.Dump(err)
	}
	country, err := frdm.Countries.FindCountryWithISO(requests.FindCountryWithISO{ISO: countryISO})
	if err != nil {
		spew.Dump(err)
	}
	countries, err := frdm.Countries.FindCountriesInContinent(requests.FindCountriesInContinent{ISO: continentISO})
	if err != nil {
		spew.Dump(err)
	}
	allCountries, err := frdm.Countries.FindAllCountries(requests.FindAllCountries{})
	if err != nil {
		spew.Dump(err)
	}
	allContinents, err := frdm.Continents.FindAllContinents(requests.FindAllContinents{})
	if err != nil {
		spew.Dump(err)
	}

	spew.Printf("Continent with ISO `%s`:\n", continentISO)
	spew.Dump(continent)

	spew.Printf("Country with ISO `%s`:\n", countryISO)
	spew.Dump(country)

	spew.Printf("%d countries in %s :\n", len(countries), continentISO)
	// spew.Dump(countries)

	spew.Printf("%d countries in total :\n", len(allCountries))
	// spew.Dump(allCountries)

	spew.Printf("%d continents in total :\n", len(allContinents))
	spew.Dump(allContinents)
}

// API is the external struct to intefacing with all service entities
type API struct {
	Continents service.Continents
	Countries  service.Countries
}

// NewAPI creates a new API to interact with
func NewAPI(continents service.Continents, countries service.Countries) *API {
	return &API{continents, countries}
}
