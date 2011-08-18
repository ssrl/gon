package signup

import . "framework/mv"
import mymy "github.com/ziutek/mymysql"
import "gaz"

type SignupController struct {
    Params   map[string]string
    Connection  *gaz.Connection
}

func (c *SignupController) Index() Model {
    col := c.Connection.DB("test").C("user")
    _, _ = col.Insert(&map[string]string{"name": c.Params["name"], "password": c.Params["email"], "email": c.Params["email"]})
    result := make(map[string]string)
    data := col.FindOne(gaz.Params{"name": c.Params["name"]}).(*mymy.Row)
	result["name"] = data.Str(1)
	result["email"] = data.Str(3)
	return Model{"name": result["name"], "email": result["email"]}
}
