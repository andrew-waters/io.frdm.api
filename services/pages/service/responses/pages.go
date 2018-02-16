package responses

import (
	"time"
)

type Page struct {
	ID        string
	Partition string
	CreatedAt time.Time
	Author    string
	URI       string
	Status    string
}
