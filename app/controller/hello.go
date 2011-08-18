package hello

import . "framework/mv"
import "app/domain/user"
import "fmt"

type HelloController struct {
    Params
    *user.UserService
}

func (c *HelloController) Index() Model {
    fmt.Printf("id = %s\n", c.Params["id"])
	return Model{"user": ""}
	// book.Get(c.Params["id"])}

	//
	// /hello/index/doc1
	// /:controller/:action/:id
	//
    // return book.findById(Params["id"])
	//
}
