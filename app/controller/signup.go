package signup

import . "framework/mv"

type SignupController struct {
  // Params
}

func (c *SignupController) Index() Model {

	return Model{"book": ""}//book.Get(c.Params["id"])}

}
