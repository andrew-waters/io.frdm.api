package mongobase

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ApplyLimit adds a limit to the query resultset
func ApplyLimit(query *mgo.Query, limit int) {
	if limit > 0 {
		query.Limit(limit)
	}
}

// ApplyOffset offsets the resultset
func ApplyOffset(query *mgo.Query, offset int) {
	if offset > 0 {
		query.Skip(offset)
	}
}

// Fields limits the fields selected by a query: https://godoc.org/labix.org/v2/mgo#Query.Select
func Fields(query *mgo.Query, fields []string) {
	if len(fields) > 0 {
		sel := bson.M{}
		for _, field := range fields {
			sel[field] = 1
		}
		query = query.Select(sel)
	}
}

// ApplySort adds sorting to a query: https://godoc.org/labix.org/v2/mgo#Query.Sort
func ApplySort(query *mgo.Query, sort map[string]int) {
	if len(sort) > 0 {
		var fields []string
		for field, order := range sort {
			if order == -1 {
				field = "-" + field
			}
			fields = append(fields, field)
		}
		query = query.Sort(fields...)
	}
}
