package requests

// FindPageByID request
type FindPageByID struct {
	Partition string
	ID        string
	Fields    []string
}
