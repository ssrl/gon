package hello

import . "framework/mv"

import "launchpad.net/gobson/bson"
import "launchpad.net/mgo"

import "app/domain/book"

var Session = new(mgo.Session)

func Index() Model {

	return Model{"book": book.Get("doc1")}

	//
	// /hello/index/doc1
	// /:controller/:action/:id
	//
    // return book.findById(Params["id"])
	//
}
