package mongobase

import mgo "gopkg.in/mgo.v2"

// Info contains detais for connecting to a MongoDB deployment
type Info struct {
	DSN string
}

// Connect to a MongoDB deployment and return a session
func (i Info) Connect() (*mgo.Session, error) {
	return mgo.Dial(i.DSN)
}
