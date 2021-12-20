# Assert

A simple package providing a function to assert two given values are deeply
equal (according to `reflect.DeepEqual`) and fail a test if they are not.

```go
package main

import (
	"testing"

	"github.com/itsliamegan/assert"
)

func TestExample(t *testing.T) {
	assert := assert.New(t)

	result := 2 + 2

	assert.Equal(result, 4)
}
```
