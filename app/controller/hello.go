package hello

import . "framework/mv"
import "app/domain/book"
import "fmt"

type HelloController struct {
    Params
    *book.BookService
}

func (c *HelloController) Index() Model {
    fmt.Printf("id = %s\n", c.Params["id"])
	return Model{"book": ""}
	// book.Get(c.Params["id"])}

	//
	// /hello/index/doc1
	// /:controller/:action/:id
	//
    // return book.findById(Params["id"])
	//
}
