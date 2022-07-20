package skiplist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipList_Add(t *testing.T) {
	skipList := NewSkipList(PromoteRate)

	skipList.Add(1)
	skipList.Add(2)
	skipList.Add(3)
	skipList.Add(4)
	skipList.Add(5)
	skipList.Add(6)
	skipList.Add(7)
}

func TestSkipList_Search(t *testing.T) {
	skipList := NewSkipList(PromoteRate)

	skipList.Add(1)
	skipList.Add(2)
	skipList.Add(3)
	skipList.Add(4)
	skipList.Add(5)
	skipList.Add(6)
	skipList.Add(7)

	node := skipList.Search(1)
	assert.Equal(t, int32(1), node.val)

	node = skipList.Search(5)
	assert.Equal(t, int32(5), node.val)

	node = skipList.Search(8)
	assert.Nil(t, node)
}
