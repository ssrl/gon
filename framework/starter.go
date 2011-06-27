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

func ToUpperFirstLetter(name string) string {
	return strings.ToUpper(string(name[0:1])) + name[1:]	
}

func FindMethod(actionMethName string, controllerType reflect.Type) (reflect.Method, bool) {
	controllerTypePointer := reflect.PtrTo(controllerType)
	var actionMeth reflect.Method
	found := false
	numMethod := controllerTypePointer.NumMethod()
	for i:=0; i < numMethod ;i++ {
        if controllerTypePointer.Method(i).Name == actionMethName {
            actionMeth = controllerTypePointer.Method(i)
            found = true
            break
        }
    }
	return actionMeth, found	
}

func RenderWithActionName(context *web.Context, ret[] reflect.Value) {
	m := ret[0].Interface().(mv.Model)
    v := ret[1].Interface().(mv.View)
    controllerName := v.String()
    context.WriteString(mustache.RenderFile("app/view/" + controllerName + "/index.m", m))
}

func RenderDefault(context *web.Context, ret[] reflect.Value, controllerName string, actionName string) {
	m := ret[0].Interface().(mv.Model)
    context.WriteString(mustache.RenderFile("app/view/" + controllerName + "/" + actionName + ".m", m))	
}

func InjectValues(context *web.Context, controllerType reflect.Type, actionMeth reflect.Method) []reflect.Value{
    conValue := reflect.New(controllerType)
    conIndirect := reflect.Indirect(conValue)

    // Inject Params
    conIndirect.FieldByName("Params").Set(reflect.ValueOf(context.Request.Params))

    // Inject beans
    for beanName,setterFunc := range bean.Registry() {
        if _, ok := controllerType.FieldByName(beanName); ok {
            if f := conIndirect.FieldByName(beanName); f.IsValid() {
                f.Set(reflect.ValueOf(setterFunc()))
            }
        }
    }

    action := actionMeth.Func
    return action.Call([]reflect.Value{conValue})
}

func Get(context *web.Context, val string) {
	controllerName, actionName := SplitControllerAndAction(val)

    if controllerType,ok := C.Controllers[controllerName]; ok {
        
        actionMethName := ToUpperFirstLetter(actionName)
        var actionMeth, found = FindMethod(actionMethName, controllerType)

        if !found { return }
		
		ret := InjectValues(context, controllerType, actionMeth);

        if len(ret) == 2 {
            RenderWithActionName(context, ret)
        } else if len(ret) == 1 {
            RenderDefault(context, ret, controllerName, actionName)
        }
    }
    return
}