package entities

import (
	"testing"
	"time"
)

func TestIsPublished(t *testing.T) {

	page := Page{
		Status: "draft",
	}

	if page.IsPublished() {
		t.Error("Drafted page should not be published")
	}

	page.Status = "live"
	if !page.IsPublished() {
		t.Error("Live page should be published")
	}

	yesterday := time.Now().Add(-24 * time.Hour)
	tomorrow := time.Now().Add(24 * time.Hour)
	dayaftertomorrow := time.Now().Add(48 * time.Hour)

	page.PublishedTime = PublishedTime{
		From: &yesterday,
		To:   &tomorrow,
	}

	if !page.IsPublished() {
		t.Error("Page should be published")
	}

	page.PublishedTime = PublishedTime{
		From: &tomorrow,
		To:   &dayaftertomorrow,
	}

	if page.IsPublished() {
		t.Error("Page should NOT be published")
	}
}
