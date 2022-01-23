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

func (assert Assert) Ok(value interface{}) {
	if value == nil {
		assert.test.Fatal("expected non-nil")
	}
}

func (assert Assert) NoErr(err error) {
	if err != nil {
		assert.test.Fatalf("expected no error, got %#+v", err)
	}
}

func (assert Assert) Equal(actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		assert.test.Fatalf("expected %#+v, got %#+v", expected, actual)
	}
}
