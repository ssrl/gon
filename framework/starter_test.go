package starter

import "github.com/bmizerany/assert"
import "testing"

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
