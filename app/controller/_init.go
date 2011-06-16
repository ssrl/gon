package controller

import "reflect"
import "app/controller/hello"

var Controllers = map[string]reflect.Value{}
var Injectables = map[string]reflect.Value{}

func registerController(s string, c interface{}) {
    v := reflect.ValueOf(c)
    Controllers[s] = v
}

func registerInjectable(s string, c interface{}) {
    v := reflect.ValueOf(c)
    Injectables[s] = v
}

func init() {
    registerController("hello/index", hello.Index)
    registerInjectable("hello.Session", hello.Session)
}
