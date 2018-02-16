package memory

import (
	"encoding/json"
	"io/ioutil"

	"github.com/andrew-waters/frdm/services/locale/core/entities"
	"github.com/andrew-waters/frdm/services/locale/core/errors"
)

type ContinentRepository struct {
	continents []entities.Continent
}

func NewContinentRepository() (*ContinentRepository, error) {

	// read all continents from disk and add them to memory
	var continents []entities.Continent
	fileData, err := fileContents("./datastore/memory/continents.json")
	if err != nil {
		return &ContinentRepository{}, err
	}
	json.Unmarshal(fileData, &continents)

	// return the repository
	return &ContinentRepository{
		continents: continents,
	}, nil
}

func fileContents(file string) ([]byte, error) {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return make([]byte, 0), err
	}
	return raw, nil
}

func (c ContinentRepository) FindByISO(ISO string) (entities.Continent, error) {

	for _, continent := range c.continents {
		if continent.ISO == ISO {
			return continent, nil
		}
	}

	return entities.Continent{}, errors.ErrContinentNotFound
}
