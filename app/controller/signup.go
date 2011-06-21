package signup

import . "framework/mv"
import "launchpad.net/mgo"
import "launchpad.net/gobson/bson"

type SignupController struct {
    Params   map[string]string
    Session  *mgo.Session
}

func (c *SignupController) Index() Model {
    defer c.Session.Close()
    col := c.Session.DB("gon").C("book")
    _ = col.Insert(&map[string]string{"_id": c.Params["email"]})
    result := make(map[string]string)
    col.Find(bson.M{"_id": c.Params["email"]}).One(result)
	return Model{"email": result["_id"]}
}
