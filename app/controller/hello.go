package hello

import . "framework/mv"
import "app/domain/book"
import "launchpad.net/mgo"

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
