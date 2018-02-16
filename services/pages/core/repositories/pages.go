package repositories

import "github.com/andrew-waters/frdm/services/pages/core/entities"

// Pages provides access to page storage
type Pages interface {
	FindByID(partition string, ID string, fields []string) (entities.Page, error)
}
