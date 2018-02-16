package entities

import "github.com/andrew-waters/frdm/services/pages/service"

// Entity represents an entity in the application
type Entity interface {
	Validate(service.Request) error
	Response() service.Response
}
