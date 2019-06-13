package selection

import (
	"../util"
	"gotest.tools/assert"
	"testing"
)

func Test_findAddressOfMinValue(t *testing.T) {
	arr := [...]int{1, 2, 3}
	address := findAddressOfMinValue(arr[:])

	// 最小的元素的地址就是arr[0]的地址
	assert.Equal(t, &arr[0], address)
}

func TestSort(t *testing.T) {
	arr := util.GenerateRandomArray(40)
	assert.Check(t, util.IsSortArray(arr) == false)

	Sort(arr)
	assert.Check(t, util.IsSortArray(arr))
}

// 1624687857 ns/op	  370264 B/op
func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := util.GenerateRandomArray(40 * 1000)
		assert.Check(b, util.IsSortArray(arr) == false)

		Sort(arr)
		assert.Check(b, util.IsSortArray(arr))
	}
}

// 6759819785 ns/op	  730648 B/op
func BenchmarkSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := util.GenerateRandomArray(80 * 1000)
		assert.Check(b, util.IsSortArray(arr) == false)

		Sort(arr)
		assert.Check(b, util.IsSortArray(arr))
	}
}
/**
可以看到，单次排序需要的时间和数组的长度不是线性关系，而是平方关系的
 */