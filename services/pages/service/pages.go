package service

import (
	"github.com/andrew-waters/frdm/services/pages/service/requests"
	"github.com/andrew-waters/frdm/services/pages/service/responses"
)

// Pages is a boundary that can do things with pages
type Pages interface {
	FindByID(requests.FindPageByID) (responses.Page, error)
}
