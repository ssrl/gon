package hello

import . "framework/mv"
import "app/domain/book"
import "launchpad.net/mgo"

var Session = new(mgo.Session)

type HelloController struct {
    Params  map[string]string
    Session *mgo.Session  
}

func (c *HelloController) Index() Model {

	return Model{"book": book.Get(c.Params["id"])}

	//
	// /hello/index/doc1
	// /:controller/:action/:id
	//
    // return book.findById(Params["id"])
	//
}
