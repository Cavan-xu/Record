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
