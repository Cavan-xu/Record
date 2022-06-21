package binarysearch

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	t.Log(Sqrt(1))
	t.Log(Sqrt(4))
	t.Log(Sqrt(10000))
	t.Log(Sqrt(2)) // 网上计算：1.414213562373，程序计算：1.414213562373095，无限接近
}
