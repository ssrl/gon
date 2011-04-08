package main

import "web"
import "reflect"
import "mustache"
import C "controller"
import "mv"

func get(ctx *web.Context, val string) {
    // a := &C.AppController{}
    // typ := reflect.Typeof(a)
    f,ok := C.Controllers[val]
    if ok {
        // av := reflect.NewValue(a)
        ret := f.Call([]reflect.Value{})
        var controllerName string = ""
        if len(ret) == 2 {
            m := ret[0].Interface().(mv.Model)
            v := ret[1].Interface().(mv.View)
            controllerName = v.String()
            ctx.WriteString(mustache.RenderFile("view/" + controllerName + "/index.m", m))
        } else if len(ret) == 1 {
            m := ret[0].Interface().(mv.Model)
            ctx.WriteString(mustache.RenderFile("view/" + val + ".m", m))
        }
    }
    return
}

func main() {
    web.Get("/(.*)", get)
    web.Run("0.0.0.0:8080")
}
