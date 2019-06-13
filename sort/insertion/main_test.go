package insertion

import (
	"fmt"
	"testing"

	"../util"
	"gotest.tools/assert"
)

func Test_fixLastIndex(t *testing.T) {
	arr := [...]int{1, 3, 4, 5, 2}
	arrS := arr[:]

	fixLastIndex(arrS)
	assert.Check(t, util.IsSortArray(arrS))
}

func TestSort(t *testing.T) {
	arr := util.GenerateRandomArray(40)
	assert.Check(t, util.IsSortArray(arr) == false)

	Sort(arr)
	fmt.Println(arr)
	assert.Check(t, util.IsSortArray(arr))
}

// 588881940 ns/op
func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := util.GenerateRandomArray(40 * 1000)
		assert.Check(b, util.IsSortArray(arr) == false)

		Sort(arr)
		assert.Check(b, util.IsSortArray(arr))
	}
}

// 2390975937 ns/op
func BenchmarkSort2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := util.GenerateRandomArray(80 * 1000)
		assert.Check(b, util.IsSortArray(arr) == false)

		Sort(arr)
		assert.Check(b, util.IsSortArray(arr))
	}
}
/**
比选择排序明显快一点
 */
