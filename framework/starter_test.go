package starter

import "github.com/bmizerany/assert"
import "testing"
import . "framework/mv"
import "reflect"


func TestSplitControllerAndAction(t *testing.T) {
    v1, v2 := splitControllerAndAction("hello/index")
    assert.Equal(t, v1, "hello")
    assert.Equal(t, v2, "index")
}

func TestSplitControllerAndActionWithDefault(t *testing.T) {
    v1, v2 := splitControllerAndAction("hello")
    assert.Equal(t, v1, "hello")
    assert.Equal(t, v2, "index")
}

func TestSplitControllerSlashAndActionWithDefault(t *testing.T) {
    v1, v2 := splitControllerAndAction("hello/")
    assert.Equal(t, v1, "hello")
    assert.Equal(t, v2, "index")
}

func TestToUpperFirstLetter(t *testing.T) {
    assert.Equal(t, toUpperFirstLetter("index"), "Index")
    assert.Equal(t, toUpperFirstLetter("indexAndIndex"), "IndexAndIndex")
    assert.Equal(t, toUpperFirstLetter(""), "")
    assert.Equal(t, toUpperFirstLetter("a"), "A")
    assert.Equal(t, toUpperFirstLetter("A"), "A")
}

type HelloController struct {
    Params
}
func (h *HelloController) Index() Model {
    return Model{"key":"value"}
}

func TestFindMethodAndInvoke(test *testing.T) {
    t := reflect.ValueOf(HelloController{}).Type()
    _, ok1 := findMethod("NotExist", t)
    assert.Equal(test, ok1, false)

    f, ok2 := findMethod("Index", t)
    assert.Equal(test, ok2, true)
    ret := f.Call([]reflect.Value{reflect.ValueOf(&HelloController{})})
    assert.Equal(test, ret[0].Interface().(Model), Model{"key":"value"})
}

type MockWebContext struct {
    result string
}
func (c *MockWebContext) WriteString(content string) {
    c.result = content
}
func (c *MockWebContext) GetParams() map[string]string {
    return map[string]string{"k":"v"}
}

func TestRenderDefault_ReturnModel(test *testing.T) {
    c := new(MockWebContext)
    ret := []reflect.Value{reflect.ValueOf(Model{"key":"test"})}    
    renderDefault(c, ret, "mock","index")
    assert.Equal(test, c.result, "mock-test")
}