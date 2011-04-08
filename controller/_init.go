package controller

import "reflect"
import "controller/hello"

var Controllers = map[string]*reflect.FuncValue{}

func register(s string, c interface{}) {
    v := reflect.NewValue(c).(*reflect.FuncValue)
    Controllers[s] = v
}

func init() {
    register("hello/index", hello.Index)
}
