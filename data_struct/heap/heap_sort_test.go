package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int32{6, 3, 8, 2, 4, 7, 9, 1}
	want := []int32{1, 2, 3, 4, 6, 7, 8, 9}

	assert.Equal(t, want, Sort(arr))
}
