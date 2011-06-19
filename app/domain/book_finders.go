package book

/*
import "framework/bean"
import "launchpad.net/gobson/bson"
import "launchpad.net/mgo"
*/

// var Template = goom.Template

func Get(id string) (result *Book) {
	return nil
	/*
	result,_ = Template.QueryOne(func()interface{}{
		c := Session.DB("gon").C("book")
		r = new(Book)
		_ = c.Find(bson.M{"_id": id}).One(r)
		return r
	}).(*Book)
	*/
}

func FindByName(name string) {

}
