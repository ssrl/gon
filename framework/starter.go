package starter

import "web"
import "reflect"
import "mustache"
import C "app/controller"
import "framework/mv"
import "strings"
import "app/conf/bean"
import "app/conf/bootstrap"

func Start() {
  bean.Init()
  bootstrap.BootStrap()
}

func Get(ctx *web.Context, val string) {
    v := strings.Split(val,"/",2)
    controllerName,actionName := v[0],v[1]
    if conType,ok := C.Controllers[controllerName]; ok {
        conTypePtr := reflect.PtrTo(conType)
        actionMethName := strings.ToUpper(string(actionName[0:1])) + actionName[1:]
        var actionMeth reflect.Method
        found := false
        for i:=0; i<conTypePtr.NumMethod();i++ {
            if conTypePtr.Method(i).Name == actionMethName {
                actionMeth = conTypePtr.Method(i)
                found = true
                break
            }
        }
        if !found { return }
        conValue := reflect.New(conType)
        reflect.Indirect(conValue).FieldByName("Params").Set(reflect.ValueOf(ctx.Request.Params))
        /*
        // NumMethod returns the number of methods in the type's method set.    
        for beanName,setterFunc := range bean.Registry() {
            if target, ok := C.Injectables[controllerName + "." + beanName]; ok {
                v := reflect.ValueOf(setterFunc())
                reflect.Indirect(target).Set(reflect.Indirect(v))
            }
        }
        */
        action := actionMeth.Func
        ret := action.Call([]reflect.Value{conValue})
        if len(ret) == 2 {
            m := ret[0].Interface().(mv.Model)
            v := ret[1].Interface().(mv.View)
            controllerName = v.String()
            ctx.WriteString(mustache.RenderFile("app/view/" + controllerName + "/index.m", m))
        } else if len(ret) == 1 {
            m := ret[0].Interface().(mv.Model)
            ctx.WriteString(mustache.RenderFile("app/view/" + controllerName + "/" + actionName + ".m", m))
        }
    }
    return
}