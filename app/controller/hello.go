package hello

import . "framework/mv"
import  "launchpad.net/gobson/bson"
import  "launchpad.net/mgo"

var Session = new(mgo.Session)

func findById(c mgo.Collection, id string) Model {
	m := Model{}
	_ = c.Find(bson.M{"_id": id}).One(&m)	
	return m
}

func Index() Model {
	book := Session.DB("gon").C("book")
	defer Session.Close()
	return findById(book, "doc1")

	//
	// /hello/index/doc1
	// /:controller/:action/:id
	//
    // return book.findById(Params["id"])
	//
}
