package util

/**
交换两个地址的值
*/
func Swap(a *int, b *int)  {
	tem := *a
	*a = *b
	*b = tem
}