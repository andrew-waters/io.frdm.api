package interactors

import (
	"github.com/andrew-waters/frdm/services/pages/core/entities"
	"github.com/andrew-waters/frdm/services/pages/core/repositories"
	"github.com/andrew-waters/frdm/services/pages/service/requests"
	"github.com/andrew-waters/frdm/services/pages/service/responses"
)

// PageInteractor contains the repository for interacting with
type PageInteractor struct {
	repository repositories.Pages
}

// NewPageInteractor returns a usable Page Interactor
func NewPageInteractor(repository repositories.Pages) *PageInteractor {
	return &PageInteractor{
		repository: repository,
	}
}

// FindByID retrieves a page with a known partition and ID
func (interactor PageInteractor) FindByID(request requests.FindPageByID) (responses.Page, error) {

	var page entities.Page

	// validate the incoming request
	if err := page.Validate(request); err != nil {
		return responses.Page{}, err
	}

	// finished validation, find the product
	page, err := interactor.repository.FindByID(request.Partition, request.ID, request.Fields)
	if err != nil {
		return responses.Page{}, err
	}

	// return a Page response from the entity
	return page.Response(), nil
}
