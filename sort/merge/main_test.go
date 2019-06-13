package merge

import (
	"fmt"
	"testing"

	"../util"
	"gotest.tools/assert"
)

func Test_middle(t *testing.T) {
	fmt.Println("3 / 2 == ", 3 / 2)
	// 1 向下取整

	arr := [...]int{1, 3, 5, 7}
	middle := len(arr) / 2

	fmt.Println("left: ", arr[:middle])
	// [1 3]

	fmt.Println("right: ", arr[middle:])
	// [5 7]
}

func Test_merge(t *testing.T) {
	arr1 := [...]int{1, 3, 5, 7}
	arr2 := [...]int{2, 4, 6, 8}

	arr := merge(arr1[:], arr2[:])
	fmt.Println(arr)
	assert.Check(t, util.IsSortArray(arr))
}

func TestSort(t *testing.T) {
	arr := util.GenerateRandomArray(20)
	assert.Check(t, util.IsSortArray(arr) == false)

	arr = Sort(arr)
	fmt.Println(arr)
	assert.Check(t, util.IsSortArray(arr))
}

// 4582226 ns/op
func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := util.GenerateRandomArray(40 * 1000)
		assert.Check(b, util.IsSortArray(arr) == false)

		arr = Sort(arr)
		assert.Check(b, util.IsSortArray(arr))
	}
}

// 9213395 ns/op
func BenchmarkSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := util.GenerateRandomArray(80 * 1000)
		assert.Check(b, util.IsSortArray(arr) == false)

		arr = Sort(arr)
		assert.Check(b, util.IsSortArray(arr))
	}
}

/**
比选择排序明显快一点
*/

