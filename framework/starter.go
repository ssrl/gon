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
        	//
        	// return only model
        	//
		 	case len(ret) == 1:
	            if m,ok := ret[0].Interface().(mv.Model); ok {
	            	mfile := "app/view/" + controllerName + "/" + actionName + ".m"
	            	ctx.WriteString(mustache.RenderFile(mfile, m))
	            } else if view,ok := ret[0].Interface().(mv.View); ok {
	            	actionName := view.String()
	            	mfile := "app/view/" + controllerName + "/" + actionName + ".m"
	            	ctx.WriteString(mustache.RenderFile(mfile, nil))
	            }
    	    
    		//
    		// return Model and View
    		//
    		case len(ret) == 2:
	            m := ret[0].Interface().(mv.Model)
	            v := ret[1].Interface().(mv.View)
	            controllerName = v.String()
	            ctx.WriteString(mustache.RenderFile("app/view/" + controllerName + "/index.m", m))    			
        }
    }
    return
}