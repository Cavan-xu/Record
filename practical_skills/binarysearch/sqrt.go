package binarysearch

/*
	使用二分法求平方根
*/

const maxDepth = 64

func Sqrt(f float64) float64 {
	if f < 0 {
		return 0
	}

	depth := 0
	sum := 1 * f
	max := float64(1)
	min := f

	if max < min {
		max = f
		min = float64(1)
	}

	for min < max {
		mid := (max + min) / 2
		if mid*mid > sum {
			max = mid
		} else {
			min = mid
		}
		depth++
		if depth > maxDepth {
			break
		}
	}

	if max*max <= sum {
		return max
	} else {
		return min
	}
}
