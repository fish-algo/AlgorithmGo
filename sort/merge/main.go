package merge

/**
归并排序的和核心就是分 + 和

分: 将一个大数组一分为二，每次只排序这一个
合: 将两个处于有序状态的小数组给合并一下
 */

/**
要求数组arr1 和 arr2都是有序数组

这里希望把它们合并成一个有序数组

创建一个新的数组，每次使用原来两个数组的最小值来填充
*/
func merge(sortArray1 []int, sortArray2 []int) []int {
	arr := make([]int, len(sortArray1) + len(sortArray2))
	currentIndex1 := 0
	currentIndex2 := 0
	for i := 0; i < len(arr); i++ {
		// 说明数组1到头了，以后只能使用数组2
		if currentIndex1 == len(sortArray1) {
			arr[i] = sortArray2[currentIndex2]
			currentIndex2++
			continue
		}
		// 说明数组2到头了，以后只能使用数组1
		if currentIndex2 == len(sortArray2) {
			arr[i] = sortArray1[currentIndex1]
			currentIndex1++
			continue
		}

		if sortArray1[currentIndex1] < sortArray2[currentIndex2] {
			arr[i] = sortArray1[currentIndex1]
			currentIndex1++
		} else {
			arr[i] = sortArray2[currentIndex2]
			currentIndex2++
		}
	}
	return arr
}

/**
遍历数组，每次都将最后一个元素排序

将数组逐步分割，需要分割logN个层级
对每个层级进行排序和merge是O(n)的复杂度

O(N * logN )
*/
func Sort(arr []int) []int  {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	leftArr := Sort(arr[:middle])
	rightArr := Sort(arr[middle:])

	return merge(leftArr, rightArr)
}
