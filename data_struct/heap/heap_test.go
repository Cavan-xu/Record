package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHeap(t *testing.T) {
	arr := []int32{6, 3, 8, 2, 4, 7, 9, 1}
	want := []int32{1, 2, 7, 3, 4, 8, 9, 6}

	heap := NewHeap(arr)
	assert.Equal(t, want, heap.Arr)
}

func TestHeap_DeleteTop(t *testing.T) {
	arr := []int32{6, 3, 8, 2, 4, 7, 9, 1}
	want := []int32{2, 3, 7, 6, 4, 8, 9}

	heap := NewHeap(arr)
	res := heap.DeleteTop()

	assert.Equal(t, int32(1), res)
	assert.Equal(t, want, heap.Arr)
}
