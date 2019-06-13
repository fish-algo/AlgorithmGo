package insertion

import "../util"

/**
插入排序
O(n * n)
 */

/**
假设数组arr除了最后一个元素，其他的元素都是出于排序状态

这样函数就是让最后一个元素处于合适的位置
 */
func fixLastIndex(arr []int)  {
	// 从后向前遍历数组
	for i := len(arr) - 1; i > 0; i-- {

		// 如果后一个元素小于前一个元素，那么需要调换位置
		if arr[i] < arr[i - 1] {
			util.Swap(&arr[i], &arr[i - 1])
		} else {
			// 否则说明已近完成了排序，直接返回
			return
		}
	}
}

/**
遍历数组，每次都将最后一个元素排序
 */
func Sort(arr []int)  {
	for i := 1; i <= len(arr); i++ {
		fixLastIndex(arr[0:i])
	}
}