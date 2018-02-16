package main

import (
	"fmt"
	"os"

	"github.com/andrew-waters/frdm/services/locale/core/interactors"
	"github.com/andrew-waters/frdm/services/locale/datastore/memory"
	"github.com/andrew-waters/frdm/services/locale/service"
	"github.com/andrew-waters/frdm/services/locale/service/requests"
	"github.com/davecgh/go-spew/spew"
)

func main() {

	// create concrete repositories
	continentRepository, err := memory.NewContinentRepository()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// create the service API
	frdm := NewAPI(
		interactors.NewContinentInteractor(continentRepository),
	)

	continent, err := frdm.Continents.FindByISO(requests.FindContinentByISO{ISO: "NA"})

	spew.Dump(err)
	spew.Dump(continent)

}

// API is the external struct to intefacing with all service entities
type API struct {
	Continents service.Continents
}

// NewAPI creates a new API to interact with
func NewAPI(continents service.Continents) *API {
	return &API{continents}
}
