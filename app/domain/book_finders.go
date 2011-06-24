/**
  This should be an auto-generated file
  similar to those in the ROO framework.
**/

package book

import "launchpad.net/gobson/bson"
import "launchpad.net/mgo"

type BookService struct {
    *mgo.Session
}

func (b *BookService) Get(id string) (result *Book) {
	c := b.DB("gon").C("book")
	defer b.Close()
	result = new(Book)
	_ = c.Find(bson.M{"_id": id}).One(result)
	return
}

func FindByName(name string) {

}
