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

// It should accept more than two variables.  
func SplitControllerAndAction(value string) (string,string) {
	controllerAndActionName := strings.Split(value,"/",2)
	controllerName := ""
    actionName  := ""

	if len(controllerAndActionName) == 2 {
        controllerName,actionName = controllerAndActionName[0],controllerAndActionName[1]
    } else if len(controllerAndActionName) == 1 {
        controllerName = controllerAndActionName[0]
        actionName = "index"
    }	

	return controllerName, actionName
}

func toUpperFirstLetter(name string) string {
	return strings.ToUpper(string(name[0:1])) + name[1:]	
}

func Get(ctx *web.Context, val string) {
	controllerName, actionName := SplitControllerAndAction(val)

    if conType,ok := C.Controllers[controllerName]; ok {
        conTypePtr := reflect.PtrTo(conType)
        actionMethName := toUpperFirstLetter(actionName)
        var actionMeth reflect.Method
        found := false
		numMethod := conTypePtr.NumMethod()
        for i:=0; i < numMethod ;i++ {
            if conTypePtr.Method(i).Name == actionMethName {
                actionMeth = conTypePtr.Method(i)
                found = true
                break
            }
        }
        if !found { return }
        conValue := reflect.New(conType)
        conIndirect := reflect.Indirect(conValue)
        conIndirect.FieldByName("Params").Set(reflect.ValueOf(ctx.Request.Params))
        
        // NumMethod returns the number of methods in the type's method set.    
        for beanName,setterFunc := range bean.Registry() {
            if f := conIndirect.FieldByName(beanName); f.IsValid() {
                f.Set(reflect.ValueOf(setterFunc()))
            }
        }

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