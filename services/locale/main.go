package main

import (
	"net/http"
	"os"

	"github.com/andrew-waters/frdm/services/locale/core/interactors"
	"github.com/andrew-waters/frdm/services/locale/datastore/memory"
	"github.com/andrew-waters/frdm/services/locale/service"
	"github.com/andrew-waters/frdm/services/locale/transport"
	"github.com/davecgh/go-spew/spew"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
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
	locales := NewLocales(
		interactors.NewContinentInteractor(continentRepository),
		interactors.NewCountryInteractor(countryRepository),
	)

	router := mux.NewRouter()
	router.Methods("GET").Path("/locales/countries").Handler(httptransport.NewServer(
		transport.MakeFindAllCountriesEndpoint(locales.Countries),
		transport.DecodeFindAllCountriesRequest,
		transport.EncodeResponse,
	))

	router.Methods("GET").Path("/locales/continents").Handler(httptransport.NewServer(
		transport.MakeFindAllContinentsEndpoint(locales.Continents),
		transport.DecodeFindAllContinentsRequest,
		transport.EncodeResponse,
	))

	router.Methods("GET").Path("/locales/continents/{iso}/countries").Handler(httptransport.NewServer(
		transport.MakeFindAllCountriesInContinentEndpoint(locales.Countries),
		transport.DecodeFindAllCountriesInContinentRequest,
		transport.EncodeResponse,
	))

	http.ListenAndServe(":8080", router)
}

// Locales is the external struct to intefacing with all service entities
type Locales struct {
	Continents service.Continents
	Countries  service.Countries
}

// NewLocales creates a new API to interact with
func NewLocales(continents service.Continents, countries service.Countries) *Locales {
	return &Locales{continents, countries}
}

// func test() {
// continentISO := "SA"
// countryISO := "GB"

// continent, err := locales.Continents.FindByISO(requests.FindContinentByISO{ISO: continentISO})
// if err != nil {
// 	spew.Dump(err)
// }
// country, err := locales.Countries.FindCountryWithISO(requests.FindCountryWithISO{ISO: countryISO})
// if err != nil {
// 	spew.Dump(err)
// }
// countries, err := locales.Countries.FindCountriesInContinent(requests.FindCountriesInContinent{ISO: continentISO})
// if err != nil {
// 	spew.Dump(err)
// }
// allCountries, err := locales.Countries.FindAllCountries(requests.FindAllCountries{})
// if err != nil {
// 	spew.Dump(err)
// }
// allContinents, err := locales.Continents.FindAllContinents(requests.FindAllContinents{})
// if err != nil {
// 	spew.Dump(err)
// }

// spew.Printf("Continent with ISO `%s`:\n", continentISO)
// spew.Dump(continent)

// spew.Printf("Country with ISO `%s`:\n", countryISO)
// spew.Dump(country)

// spew.Printf("%d countries in %s :\n", len(countries), continentISO)
// // spew.Dump(countries)

// spew.Printf("%d countries in total :\n", len(allCountries))
// // spew.Dump(allCountries)

// spew.Printf("%d continents in total :\n", len(allContinents))
// spew.Dump(allContinents)
// }
