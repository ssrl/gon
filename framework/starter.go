package starter

import "web"
import "reflect"
import "mustache"
import C "app/controller"
import "framework/mv"
import "framework/gon"
import "strings"
import "app/conf/bean"
import "app/conf/bootstrap"

const APP_VIEW_PATH = "app/view/"

func Start() {
    bean.Initialize()
    bootstrap.BootStrap(bean.DefaultAppContext)
}

func Get(context *web.Context, val string) {
    internalGet(context, val)
}

func internalGet(context gon.WebContext, val string) {
    if( len(val) == 0) {
        renderRoot(context)
        return
    }

    controllerName, actionName := splitControllerAndAction(val)

    if controllerType, ok := C.Controllers[controllerName]; ok {

        actionMethName := toUpperFirstLetter(actionName)
        if action, found := findMethod(actionMethName, controllerType); found {
            conValue := instantiateAndInjectController(context, controllerType);
            ret := action.Call([]reflect.Value{conValue})
            if len(ret) == 2 {
                // return Model and View
                renderWithActionName(context, ret, controllerName)
            } else if len(ret) == 1 {
                // return Model or View
                renderDefault(context, ret, controllerName, actionName)
            }
        }

    }
    return
}

func splitControllerAndAction(value string) (string,string) {
    controllerAndActionName := strings.Split(value,"/")
    controllerName := ""
    actionName := ""

    if len(controllerAndActionName) == 2 {
        controllerName,actionName = controllerAndActionName[0],controllerAndActionName[1]
        if actionName == "" {
            actionName = "index"
        }
    } else if len(controllerAndActionName) == 1 {
        controllerName = controllerAndActionName[0]
        actionName = "index"
    }

    return controllerName, actionName
}

func toUpperFirstLetter(name string) string {
    if name == "" {
        return ""
    } else if len(name) == 1 {
        return strings.ToUpper(name)
    }
    return strings.ToUpper(string(name[0:1])) + name[1:]
}

func findMethod(actionMethName string, controllerType reflect.Type) (reflect.Value, bool) {
    controllerTypePtr := reflect.PtrTo(controllerType)
    var actionMeth reflect.Method
    found := false
    numMethod := controllerTypePtr.NumMethod()
    for i := 0; i < numMethod; i++ {
        if controllerTypePtr.Method(i).Name == actionMethName {
            actionMeth = controllerTypePtr.Method(i)
            found = true
            break
        }
    }
    return actionMeth.Func, found
}

func renderWithActionName(context gon.WebContext, ret []reflect.Value, controllerName string) {
    model := ret[0].Interface().(mv.Model)
    view  := ret[1].Interface().(mv.View )
    actionName := view.String()
    context.WriteString(mustache.RenderFile(APP_VIEW_PATH + controllerName + "/" + actionName + ".m", model))
}

func renderDefault(context gon.WebContext, ret []reflect.Value, controllerName string, actionName string) {
    if model,ok := ret[0].Interface().(mv.Model); ok {
        context.WriteString(mustache.RenderFile(APP_VIEW_PATH + controllerName + "/" + actionName + ".m", model))
    } else if view,ok := ret[0].Interface().(mv.View); ok {
        actionName = view.String()
        context.WriteString(mustache.RenderFile(APP_VIEW_PATH + controllerName + "/" + actionName + ".m"))
    }
}

func renderRoot(context gon.WebContext){
    context.WriteString(mustache.RenderFile(APP_VIEW_PATH + "main.m"))
}

func instantiateAndInjectController(context gon.WebContext, controllerType reflect.Type) reflect.Value {
    // Instantiate a controller
    conValue := reflect.New(controllerType)
    conIndirect := reflect.Indirect(conValue)

    // Inject Params
    conIndirect.FieldByName("Params").Set(reflect.ValueOf(context.GetParams()))

    //
    // Inject beans
    // This loop tends to be slow. We should loop over field names and look-up a bean.
    //
    for beanName,setterFunc := range bean.Registry() {
        if _, ok := controllerType.FieldByName(beanName); ok {
            if field := conIndirect.FieldByName(beanName); field.IsValid() {
                field.Set(reflect.ValueOf(setterFunc()))
            }
        }
    }
    return conValue
}

