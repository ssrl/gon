package hello

import . "framework/mv"
import "app/domain/book"
import "launchpad.net/mgo"

import "fmt"

var Session = new(mgo.Session)

type HelloController struct {
    Params  map[string]string
    Session *mgo.Session  
}

func (c *HelloController) Index() Model {
    fmt.Printf("%s\n", c.Params)
	return Model{"book": book.Get(c.Params["id"])}

	//
	// /hello/index/doc1
	// /:controller/:action/:id
	//
    // return book.findById(Params["id"])
	//
}
