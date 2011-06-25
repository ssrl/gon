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

func FindMethod(actionMethName string, conType reflect.Type) (reflect.Method, bool) {
	conTypePtr := reflect.PtrTo(conType)
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
	return actionMeth, found	
}

func RenderWithActionName(ctx *web.Context, ret[] reflect.Value) {
	m := ret[0].Interface().(mv.Model)
    v := ret[1].Interface().(mv.View)
    controllerName := v.String()
    ctx.WriteString(mustache.RenderFile("app/view/" + controllerName + "/index.m", m))
}

func RenderDefault(ctx *web.Context, ret[] reflect.Value, controllerName string, actionName string) {
	m := ret[0].Interface().(mv.Model)
    ctx.WriteString(mustache.RenderFile("app/view/" + controllerName + "/" + actionName + ".m", m))	
}

func SetFunctions(ctx *web.Context, conType reflect.Type, actionMeth reflect.Method) []reflect.Value{
    conValue := reflect.New(conType)
    conIndirect := reflect.Indirect(conValue)

    // Inject Params
    conIndirect.FieldByName("Params").Set(reflect.ValueOf(ctx.Request.Params))

    // Inject beans
    for beanName,setterFunc := range bean.Registry() {
        if _, ok := conType.FieldByName(beanName); ok {
            if f := conIndirect.FieldByName(beanName); f.IsValid() {
                f.Set(reflect.ValueOf(setterFunc()))
            }
        }
    }

    action := actionMeth.Func
    return action.Call([]reflect.Value{conValue})
}

func Get(ctx *web.Context, val string) {
	controllerName, actionName := SplitControllerAndAction(val)

    if conType,ok := C.Controllers[controllerName]; ok {
        
        actionMethName := ToUpperFirstLetter(actionName)
        var actionMeth, found = FindMethod(actionMethName, conType)

        if !found { return }
		
		ret := SetFunctions(ctx, conType, actionMeth);
		
        if !found { return }

        if len(ret) == 2 {
            RenderWithActionName(ctx, ret)
        } else if len(ret) == 1 {
            RenderDefault(ctx, ret, controllerName, actionName)
        }
    }
    return
}