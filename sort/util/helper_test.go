package util

import (
	"gotest.tools/assert"
	"testing"
)

func Test_swap(t *testing.T) {
	a := 1
	b := 2
	Swap(&a, &b)
	assert.Equal(t, 2, a)
	assert.Equal(t, 1, b)
}
