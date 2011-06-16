package bean

import  "launchpad.net/mgo"

var Registry map[string]func()interface{} = map[string]func()interface{}{}

func bean(name string, f func()interface{}) {
	Registry[name] = f
}

func init() {
	bean("Session", func()interface{} {
		session, _ := mgo.Mongo("127.0.0.1")
		session.SetMode(mgo.Monotonic, true)
		return session
	})
}