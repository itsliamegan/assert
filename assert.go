package assert

import (
	"reflect"
	"testing"
)

type Assert struct {
	test *testing.T
}

func New(test *testing.T) Assert {
	return Assert{test: test}
}

func (assert Assert) Equal(expected, actual interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		assert.test.Errorf("expected: %#+v; got: %#+v", expected, actual)
	}
}
