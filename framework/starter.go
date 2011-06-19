package starter

import "web"
import "reflect"
import "mustache"
import C "app/controller"
import "framework/mv"
import "framework/bean"
import "strings"

func Post(ctx *web.Context, val string) {
	Get(ctx, val)
}

func Get(ctx *web.Context, val string) {
    v := strings.Split(val,"/",2)
    controllerName,actionName := v[0],v[1]
    if action,ok := C.Controllers[controllerName+"/"+actionName]; ok {
        for beanName,setterFunc := range bean.Registry {
            if target, ok := C.Injectables[controllerName + "." + beanName]; ok {
                v := reflect.ValueOf(setterFunc())
                reflect.Indirect(target).Set(reflect.Indirect(v))
            }
        }
        ret := action.Call([]reflect.Value{})
        switch {
		 	case len(ret) == 1:
	            m := ret[0].Interface().(mv.Model)
    	        ctx.WriteString(mustache.RenderFile("app/view/" + controllerName + "/" + actionName + ".m", m))		 		
    		case len(ret) == 2:
	            m := ret[0].Interface().(mv.Model)
	            v := ret[1].Interface().(mv.View)
	            controllerName = v.String()
	            ctx.WriteString(mustache.RenderFile("app/view/" + controllerName + "/index.m", m))    			
        }
    }
    return
}