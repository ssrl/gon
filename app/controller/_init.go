package controller

import "reflect"
import "app/controller/hello"
import "app/controller/signup"

var Controllers = map[string]reflect.Type{}
// var Injectables = map[string]reflect.Value{}

func registerController(s string, c interface{}) {
    // v := reflect.ValueOf(c)
    t := reflect.Indirect(reflect.ValueOf(c)).Type()
    Controllers[s] = t
}

/*
func registerInjectable(s string, c interface{}) {
    v := reflect.ValueOf(c)
    Injectables[s] = v
}
*/

func init() {
    registerController("hello", new(hello.HelloController))
    registerController("signup", new(signup.SignupController))
    // registerInjectable("hello.Session", hello.Session)
}
