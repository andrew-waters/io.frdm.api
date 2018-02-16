package service

import (
	"github.com/andrew-waters/frdm/services/locale/service/requests"
	"github.com/andrew-waters/frdm/services/locale/service/responses"
)

// Continents is a boundary that can do things with continents
type Continents interface {
	FindAllContinents(requests.FindAllContinents) ([]responses.Continent, error)
	FindByISO(requests.FindContinentByISO) (responses.Continent, error)
}
