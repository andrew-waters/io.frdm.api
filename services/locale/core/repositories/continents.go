package repositories

import "github.com/andrew-waters/frdm/services/locale/core/entities"

// Continents provides access to continent storage
type Continents interface {
	FindAllContinents() ([]entities.Continent, error)
	FindByISO(ISO string) (entities.Continent, error)
}
