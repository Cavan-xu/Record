package heap

/*
	二叉堆
		len(arr)/2 - 1 求到的是完全二叉树最后一个非叶子结点的下标
		解答：
			①堆的最后一个非叶子节点若只有左孩子
			②堆的最后一个非叶子节点有左右两个孩子

			完全二叉树的性质之一是：如果节点序号为i，在它的左孩子序号为2*i+1，右孩子序号为2*i+2。
			对于①左孩子的序号为n-1，则n-1=2*i-1，推出i=n/2-1；
			对于②左孩子的序号为n-2，在n-2=2*i-1，推出i=(n-1)/2-1；右孩子的序号为n-1，则n-1=2*i+2，推出i=(n-1)/2-1；
			很显然，当完全二叉树最后一个节点是其父节点的左孩子时，树的节点数为偶数；当完全二叉树最后一个节点是其父节点的右孩子时，树的节点数为奇数。
			根据语法的特征，整数除不尽时向下取整，则若n为奇数时(n-1)/2-1=n/2-1。
*/

type Heap struct {
	Arr []int32
}

func NewHeap(arr []int32) *Heap {
	heap := &Heap{}
	heap.build(arr)

	return heap
}

// 删除顶部元素
func (h *Heap) DeleteTop() int32 {
	if len(h.Arr) == 0 {
		return -1
	}

	res := h.Arr[0]

	if len(h.Arr) == 1 {
		h.Arr = []int32{}
		return res
	}

	newArr := make([]int32, len(h.Arr)-1)
	copy(newArr, h.Arr)
	newArr[0] = h.Arr[len(h.Arr)-1]

	h.downAdjust(newArr, 0, len(newArr))
	h.Arr = newArr

	return res
}

// 默认小顶堆
func (h *Heap) build(arr []int32) {
	if len(arr) == 0 {
		return
	}

	for i := len(arr)/2 - 1; i >= 0; i-- {
		h.downAdjust(arr, i, len(arr))
	}

	h.Arr = arr
}

// 向下调整
func (h *Heap) downAdjust(arr []int32, parentIndex int, length int) {
	temp := arr[parentIndex]
	childIndex := parentIndex*2 + 1

	for childIndex < length {
		if childIndex+1 < length && arr[childIndex+1] < arr[childIndex] {
			childIndex += 1
		}
		if temp < arr[childIndex] {
			break
		}
		arr[parentIndex] = arr[childIndex]
		parentIndex = childIndex
		childIndex = 2*parentIndex + 1
	}

	arr[parentIndex] = temp
}
