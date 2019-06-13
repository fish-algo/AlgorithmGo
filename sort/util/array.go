package util

import (
	"math/rand"
)

/**
判断一个数组是不是处于排序状态(从小到大)
*/
func IsSortArray(arr []int) bool {
	// 遍历数组，确认后面的元素比前面的小
	for i := 1; i < len(arr); i++{
		if arr[i] < arr[i - 1] {
			return false
		}
	}
	return true
}

/**
获取一个随机int数组的切片
 */
func GenerateRandomArray(length int) []int {
	byteArr := make([]byte, length)
	rand.Read(byteArr)

	intArr := make([]int, length)
	for i, v := range byteArr {
		intArr[i] = int(v)
	}

	return intArr
}

