package starter

import "github.com/bmizerany/assert"
import "testing"

func TestSplitControllerAndAction(t *testing.T) {
    v1, v2 := splitControllerAndAction("a/b")
    assert.Equal(t, v1, "a")
    assert.Equal(t, v2, "b")
}

func TestSplitControllerAndActionWithDefault(t *testing.T) {
    v1, v2 := splitControllerAndAction("a")
    assert.Equal(t, v1, "a")
    assert.Equal(t, v2, "index")
}