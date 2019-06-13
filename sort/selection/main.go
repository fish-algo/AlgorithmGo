package selection

import "../util"

/**
寻找到切片中最小值元素所在的位置
O(N) 时间复杂度

初始默认最小值的索引为0
然后遍历这个数组，看看有没有更小的
 */
func findAddressOfMinValue(arr []int) *int {
	minIndex := 0
	for i, v := range arr {
		if v < arr[minIndex] {
			minIndex = i
		}
	}
	return &arr[minIndex]
}

/**
选择排序：
  每次选择第i个元素以及它后面的数组的最小值
  将这个最小值和第i个元素交换

O(n * n) 时间复杂度
 */
func Sort(arr []int)  {
	for i, _ := range arr{
		minAddress := findAddressOfMinValue(arr[i:])
		util.Swap(minAddress, &arr[i])
	}
}