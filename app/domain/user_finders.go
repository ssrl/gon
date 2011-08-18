package user

import "gaz"
import mymy "github.com/ziutek/mymysql"

type UserService struct {
    *gaz.Connection
    // EntityManager
}

func (u *UserService) Get(id string) (result *User) {
	c := u.DB("test").C("User")
	result = new(User)
	data := c.FindOne(gaz.Params{"id":id}).(*mymy.Row)
	result.Id = data.Int(0)
	result.Name = data.Str(1)
	result.Password = data.Str(2)
	result.Email = data.Str(3)
	return
}