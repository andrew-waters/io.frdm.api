package entities

import (
	"time"

	"github.com/andrew-waters/frdm/services/pages/service"
	"github.com/andrew-waters/frdm/services/pages/service/responses"
)

var validStatuses = []string{
	"live",
	"draft",
	"hidden",
}

// Page represents a single page
type Page struct {
	ID            string `json:"id"`
	Partition     string `json:"partition"`
	Names         PageNames
	CreatedAt     time.Time
	Author        string
	URI           string
	Status        string
	PublishedTime PublishedTime
}

// Response returns a responses.Page for a Page
func (p Page) Response() responses.Page {
	return responses.Page{
		ID:        p.ID,
		Partition: p.Partition,
		CreatedAt: p.CreatedAt,
		Author:    p.Author,
		URI:       p.URI,
		Status:    p.Status,
	}
}

// Validate an incoming service request
func (p Page) Validate(service.Request) error {

	return nil
}

// PageNames contains the name of a page in different contexts
type PageNames struct {
	Title      string
	Slug       string
	Breadcrumb string
	Navigation string
}

// PublishedTime represents when a page should be published FROM and TO
type PublishedTime struct {
	From *time.Time
	To   *time.Time
}

// IsPublished tells us if the page is published based on it's status and PublishTime
func (p Page) IsPublished() bool {

	// the page is draft it cannot be published
	if p.Status == "draft" {
		return false
	}

	// published FROM is in the future
	if p.PublishedTime.From != nil && p.PublishedTime.From.After(time.Now()) {
		return false
	}

	// published TO is in the past
	if p.PublishedTime.To != nil && p.PublishedTime.To.Before(time.Now()) {
		return false
	}

	return true
}
