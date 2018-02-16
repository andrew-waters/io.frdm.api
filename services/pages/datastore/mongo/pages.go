package mongo

import (
	"github.com/andrew-waters/frdm/databases/mongobase"
	"github.com/andrew-waters/frdm/services/pages/core/entities"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// PageRepository satisfies the repositories.pages interface
type PageRepository struct {
	database   string
	collection string
	session    *mgo.Session
}

// NewPageRepository conveniently returns a PageRepository
func NewPageRepository(database string, collection string, session *mgo.Session) PageRepository {
	return PageRepository{
		database:   database,
		collection: collection,
		session:    session,
	}
}

// FindByID attempts to return a product with a known ID
func (mongo PageRepository) FindByID(partition string, ID string, fields []string) (entities.Page, error) {
	session := mongo.session.Copy()
	defer session.Close()

	collection := session.DB(mongo.database).C(mongo.collection)

	product := entities.Page{}

	query := collection.Find(bson.M{"partition": partition, "id": ID})

	mongobase.Fields(query, fields)

	err := query.One(&product)

	return product, err
}
