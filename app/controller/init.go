package controller

import "reflect"
import "app/controller/hello"

var Controllers = map[string]reflect.Value{}

func register(s string, c interface{}) {
    v := reflect.ValueOf(c)
    Controllers[s] = v
}

func init() {
    register("hello/index", hello.Index)
}
