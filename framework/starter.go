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
    bean.Initialize()
    bootstrap.BootStrap()
}

func splitControllerAndAction(value string) (string,string) {
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

func findMethod(actionMethName string, controllerType reflect.Type) (reflect.Value, bool) {
    controllerTypePtr := reflect.PtrTo(controllerType)
    var actionMeth reflect.Method
    found := false
    numMethod := controllerTypePtr.NumMethod()
    for i:=0; i < numMethod; i++ {
        if controllerTypePtr.Method(i).Name == actionMethName {
            actionMeth = controllerTypePtr.Method(i)
            found = true
            break
        }
    }
    return actionMeth.Func, found
}

func renderWithActionName(context *web.Context, ret[] reflect.Value) {
    m := ret[0].Interface().(mv.Model)
    v := ret[1].Interface().(mv.View)
    controllerName := v.String()
    context.WriteString(mustache.RenderFile("app/view/" + controllerName + "/index.m", m))
}

func renderDefault(context *web.Context, ret[] reflect.Value, controllerName string, actionName string) {
    m := ret[0].Interface().(mv.Model)
    context.WriteString(mustache.RenderFile("app/view/" + controllerName + "/" + actionName + ".m", m))
}

func RenderRoot(context *web.Context){
	context.WriteString(mustache.RenderFile("app/view/main.m"))
}

func instantiateAndInjectController(context *web.Context, controllerType reflect.Type) reflect.Value {
    // Instantiate a controller
    conValue := reflect.New(controllerType)
    conIndirect := reflect.Indirect(conValue)

    // Inject Params
    conIndirect.FieldByName("Params").Set(reflect.ValueOf(context.Request.Params))

    // Inject beans
    for beanName,setterFunc := range bean.Registry() {
        if _, ok := controllerType.FieldByName(beanName); ok {
            if field := conIndirect.FieldByName(beanName); field.IsValid() {
                field.Set(reflect.ValueOf(setterFunc()))
            }
        }
    }
    return conValue
}

func Get(context *web.Context, val string) {
	if( val == "") {
		RenderRoot(context)
		return
	}
    controllerName, actionName := splitControllerAndAction(val)

    if controllerType, ok := C.Controllers[controllerName]; ok {

        actionMethName := toUpperFirstLetter(actionName)
        if action, found := findMethod(actionMethName, controllerType); found {
            conValue := instantiateAndInjectController(context, controllerType);
            ret := action.Call([]reflect.Value{conValue})
            if len(ret) == 2 {
                renderWithActionName(context, ret)
            } else if len(ret) == 1 {
                renderDefault(context, ret, controllerName, actionName)
            }
        }

    }
    return
}

