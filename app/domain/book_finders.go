package book

import "app/conf/bean"

import "launchpad.net/gobson/bson"
import "launchpad.net/mgo"

func Get(id string) (result *Book) {
	session := bean.GetBean("Session").(*mgo.Session)
	c := session.DB("gon").C("book")
	defer session.Close()
	result = new(Book)
	_ = c.Find(bson.M{"_id": id}).One(result)
	return
}

func FindByName(name string) {

}
