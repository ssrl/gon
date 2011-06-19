package goom

import "launchpad.net/mgo"

type MongoTemplate struct {
	Session *mgo.Session
}

func (t *MongoTemplate) Get(id string) interface{} {
	// t.Session
	return nil
}