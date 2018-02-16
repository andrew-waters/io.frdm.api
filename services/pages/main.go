package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"github.com/andrew-waters/frdm/databases/mongobase"
	"github.com/andrew-waters/frdm/services/pages/core/interactors"
	"github.com/andrew-waters/frdm/services/pages/datastore/mongo"
	"github.com/andrew-waters/frdm/services/pages/service"
	"github.com/andrew-waters/frdm/services/pages/service/requests"
)

func main() {

	// @todo get DSN from ENV
	mongoConnInfo := mongobase.Info{
		DSN: "mongodb://XXXX:27017",
	}

	// create the database connection
	session, err := mongoConnInfo.Connect()
	if err != nil {
		fmt.Println("Couldn't create database session: ", err.Error())
	}

	// create concrete repositories
	pageRepository := mongo.NewPageRepository("pages", "pages", session)

	// create the service API
	frdm := NewAPI(
		interactors.NewPageInteractor(pageRepository),
	)

	page, err := frdm.Pages.FindByID(requests.FindPageByID{
		"1",
		"2",
		nil,
	})

	spew.Dump(err)
	spew.Dump(page)
}

// API is the external struct to intefacing with all service entities
type API struct {
	Pages service.Pages
}

// NewAPI creates a new API to interact with
func NewAPI(pages service.Pages) *API {
	return &API{pages}
}
