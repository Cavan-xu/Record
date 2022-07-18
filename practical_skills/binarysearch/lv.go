package binarysearch

/*
	通过玩家的经验总值计算出玩家的对应等级，有两种方式：
		1、逐减法：从给定的经验数组下标 0 开始递减，直到经验值小于0
			优点：实现较为简单且逻辑清晰
			缺点：如果玩家可达的等级较高的话，时间复杂度是 O(n)
		2、二分法：使用二分找到数组中第一个比该玩家经验值大的数，对应的下标+1即为玩家等级
			优点：时间复杂度 O(log(n))
			缺点：没有 bug 的二分实现并不简单！
*/

// 0~50 为等级1，51～100 为等级二
var ExpSlice = []int{50, 100, 150, 200, 250, 300, 350, 400, 450, 500, 550, 600, 650, 700, 750, 800, 850, 900, 950, 1000}

// 同样 0~49 为等级一， 50~99 为等级二，注意，下标0只占位
var ExpSliceV2 = []int{-1, 50, 100, 150, 200, 250, 300, 350, 400, 450, 500, 550, 600, 650, 700, 750, 800, 850, 900, 950, 1000}

// 使用逐减法求得玩家经验对应的等级
func UseSub(exp int) int {
	for i, v := range ExpSlice {
		if exp <= v {
			return i + 1
		}
	}

	return len(ExpSlice)
}

// 使用二分法求玩家等级
func UseBinarySearch(exp int) int {
	min := 0
	max := len(ExpSlice) - 1

	for min <= max {
		mid := (min + max) >> 1
		if ExpSlice[mid] >= exp {
			if mid == 0 || ExpSlice[mid-1] < exp {
				return mid + 1
			}
			max = mid - 1
		} else {
			min = mid + 1
		}
	}

	if ExpSlice[len(ExpSlice)-1] < exp {
		return len(ExpSlice)
	}

	return -1
}

// v2版本
func UseBinarySearchV2(exp int) int {
	min := 1
	max := len(ExpSliceV2) - 1

	for min < max-1 {
		mid := (min + max) >> 1
		if exp >= ExpSliceV2[mid] {
			min = mid
		} else {
			max = mid
		}
	}

	if ExpSliceV2[min] <= exp {
		return min + 1
	} else {
		return min
	}
}
