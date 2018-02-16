package interactors

import (
	"github.com/andrew-waters/frdm/services/locale/core/entities"
	"github.com/andrew-waters/frdm/services/locale/core/repositories"
	"github.com/andrew-waters/frdm/services/locale/service/requests"
	"github.com/andrew-waters/frdm/services/locale/service/responses"
)

// ContinentInteractor contains the repository for interacting with continents
type ContinentInteractor struct {
	repository repositories.Continents
}

// NewContinentInteractor returns a usable ContinentInteractor
func NewContinentInteractor(repository repositories.Continents) *ContinentInteractor {
	return &ContinentInteractor{
		repository: repository,
	}
}

// FindByISO retrieves a continent by ISO
func (interactor ContinentInteractor) FindByISO(request requests.FindContinentByISO) (responses.Continent, error) {

	var continent entities.Continent

	// validate the incoming request
	if err := continent.Validate(request); err != nil {
		return responses.Continent{}, err
	}

	// finished validation, find the continent
	continent, err := interactor.repository.FindByISO(request.ISO)
	if err != nil {
		return responses.Continent{}, err
	}

	// return a Page response from the entity
	return continent.Response(), nil
}

func (interactor ContinentInteractor) FindAllContinents(request requests.FindAllContinents) ([]responses.Continent, error) {

	var continent entities.Continent
	var continents []responses.Continent

	// validate the incoming request
	if err := continent.Validate(request); err != nil {
		return continents, err
	}

	// finished validation, find the continents
	matchedContinents, err := interactor.repository.FindAllContinents()
	if err != nil {
		return continents, err
	}

	for _, matchedContinent := range matchedContinents {
		continents = append(continents, matchedContinent.Response())
	}

	// return a Page response from the entity
	return continents, nil
}
