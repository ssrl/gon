package starter

import . "framework/mv"
import "github.com/bmizerany/assert"
import "testing"
import "reflect"

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

func TestRenderDefault_ReturnView(test *testing.T) {
    c := new(MockWebContext)
    ret := []reflect.Value{reflect.ValueOf(View("index"))}
    renderDefault(c, ret, "mock","non-index")
    assert.Equal(test, c.result, "mock-")
}