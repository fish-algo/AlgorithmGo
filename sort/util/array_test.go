package util

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func TestIsSortArray(t *testing.T) {
	// 这个数组应该是已经排序的了
	arr := [...]int{1, 2, 3}
	assert.Check(t, IsSortArray(arr[:]))

	// 是数组不再有序
	arr[0] = 4
	assert.Check(t, IsSortArray(arr[:]) == false)
}

func TestGenerateRandomArray(t *testing.T) {
	arr := GenerateRandomArray(44)
	fmt.Println(arr)
}
