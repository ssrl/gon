package book

var Session = new(mgo.Session)

func Get(id string) (result *Book) {
	c := Session.DB("gon").C("book")
	defer Session.Close()
	result := new(Book)
	_ = c.Find(bson.M{"_id": id}).One(result)
	return
}

func FindByName(name string) {

}

