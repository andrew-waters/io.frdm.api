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

	router.Methods("GET").Path("/locale/continents").Handler(httptransport.NewServer(
		transport.MakeFindAllContinentsEndpoint(locales.Continents),
		transport.DecodeFindAllContinentsRequest,
		transport.EncodeResponse,
	))

	router.Methods("GET").Path("/locale/continents/{iso}").Handler(httptransport.NewServer(
		transport.MakeFindContinentWithISO(locales.Continents),
		transport.DecodeFindContinentWithISORequest,
		transport.EncodeResponse,
	))

	router.Methods("GET").Path("/locale/continents/{iso}/countries").Handler(httptransport.NewServer(
		transport.MakeFindAllCountriesInContinentEndpoint(locales.Countries),
		transport.DecodeFindAllCountriesInContinentRequest,
		transport.EncodeResponse,
	))

	router.Methods("GET").Path("/locale/countries").Handler(httptransport.NewServer(
		transport.MakeFindAllCountriesEndpoint(locales.Countries),
		transport.DecodeFindAllCountriesRequest,
		transport.EncodeResponse,
	))

	router.Methods("GET").Path("/locale/countries/{iso}").Handler(httptransport.NewServer(
		transport.MakeFindCountryByISOEndpoint(locales.Countries),
		transport.DecodeFindCountryByISORequest,
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
