package heap

func Sort(arr []int32) []int32 {
	heap := NewHeap(arr)
	count := len(heap.Arr)
	res := make([]int32, 0, len(heap.Arr))

	for i := 0; i < count; i++ {
		res = append(res, heap.DeleteTop())
	}

	return res
}

func LocalSort(arr []int32) []int32 {
	heap := NewHeap(arr)
	newArr := make([]int32, len(heap.Arr))
	copy(newArr, heap.Arr)

	for i := 0; i < len(newArr); i++ {
		// 收尾交换
		val := newArr[0]
		newArr[0] = newArr[len(newArr)-i-1]
		newArr[len(newArr)-i-1] = val

		heap.downAdjust(newArr, 0, len(newArr)-i-1)
	}

	return newArr
}
