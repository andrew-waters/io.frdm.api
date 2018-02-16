package repositories

import "github.com/andrew-waters/frdm/services/locale/core/entities"

// Continents provides access to continent storage
type Continents interface {
	FindByISO(ISO string) (entities.Continent, error)
}
