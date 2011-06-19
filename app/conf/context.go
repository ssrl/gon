package bean

import "launchpad.net/mgo"
import . "framework/bean"

func init() {
	Bean("Session", func()interface{} {
		session, _ := mgo.Mongo("127.0.0.1")
		session.SetMode(mgo.Monotonic, true)
		return session
	})
	/*
	bean("Template", func()interface() {
		return &goom.MongoTemplate{
			session: ref("Session"),
		}
	})
	*/
}