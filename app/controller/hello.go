package hello

import . "framework/mv"
import  "launchpad.net/gobson/bson"
import  "launchpad.net/mgo"

func findById(c mgo.Collection, id string) Model {
	m := Model{}
	_ = c.Find(bson.M{"_id": id}).One(&m)	
	return m
}

func Index() Model {
	session, _ := mgo.Mongo("127.0.0.1")
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()
	book := session.DB("gon").C("book")
    return findById(book, "doc1")
}
